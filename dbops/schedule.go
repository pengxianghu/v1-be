package dbops

import (
	"database/sql"
	"log"
	"time"
	
	_ "github.com/go-sql-driver/mysql"
	"github.com/pengxianghu/v1-be/defs"
)

func AddSchedule(schedule *defs.Schedule) error {
	stmtIns, err := dbConn.Prepare("INSERT INTO schedule (`user_id`, `topic`, `content`) VALUES(?, ?, ?)")
	if err != nil {
		log.Printf("add schedule dbConn prepare error: %s\n", err)
		return err
	}

	_, err = stmtIns.Exec(schedule.UserId, schedule.Topic, schedule.Content)
	if err != nil {
		log.Printf("add schedule stmtIns exec error: %s\n", err)
		return err
	}

	defer stmtIns.Close()

	return nil
}

func GetScheduleByUser(u_id string) ([]*defs.Schedule, error) {
	scheduleList := make([]*defs.Schedule, 0)
	stmtOuts, err := dbConn.Prepare("SELECT `id`, `topic`,`content`,`created_at`, `status`, `active` FROM schedule WHERE `user_id` = ? AND `active` = 1 ORDER BY `created_at` DESC")
	if err != nil {
		log.Printf("get Schedule By UserId stmtOuts error: %s\n", err)
		return scheduleList, err
	}
	schedule := &defs.Schedule{}

	rows, err := stmtOuts.Query(u_id)
	// 不为空且不是no rows
	if err != nil && err != sql.ErrNoRows {
		return scheduleList, err
	}
	for rows.Next() {
		schedule = &defs.Schedule{}
		var t time.Time
		err = rows.Scan(&schedule.Id, &schedule.Topic, &schedule.Content, &t, &schedule.Status, &schedule.Active)
		schedule.CreatedAt = t.Format("2006-01-02 15:04:05")
		if err != nil {
			continue
		}
		schedule.UserId = u_id
		scheduleList = append(scheduleList, schedule)
	}

	defer stmtOuts.Close()

	return scheduleList, nil
}

func DeleteScheduleById(s_id int) error {
	stmtIns, err := dbConn.Prepare("UPDATE schedule SET `active` = 0 WHERE `id` = ?")
	if err != nil {
		log.Printf("delete schedule by id stmtIns error: %v", err)
		return err
	}
	_, err = stmtIns.Exec(s_id)
	if err != nil {
		log.Printf("delete schedule by id stmtIns exec error: %v", err)
		return err
	}

	defer stmtIns.Close()

	return nil
}