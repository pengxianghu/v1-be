/*
 * @Description:
 * @Autor: pengxianghu
 * @Date: 2019-08-11 09:25:00
 * @LastEditTime: 2019-08-17 20:38:06
 */
package dbops

import (
	"database/sql"
	"log"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/pengxianghu/v1-be/defs"
)

func GetScheduleById(id int) (*defs.Schedule, error) {
	schedule := &defs.Schedule{}
	stmtOuts, err := dbConn.Prepare("SELECT `id`, `user_id`, `topic`, `content`, `created_at`, `status`, `active` FROM schedule WHERE `id` = ?")
	if err != nil {
		log.Printf("get schedule by id err: %+v", err)
		return schedule, err
	}

	row := stmtOuts.QueryRow(id)
	var t time.Time
	err = row.Scan(&schedule.Id, &schedule.UserId, &schedule.Topic, &schedule.Content, &t, &schedule.Status, &schedule.Active)
	schedule.CreatedAt = t.Format("2006-01-02 15:04:05")
	if err != nil {
		log.Printf("get schedule by id scan err: %+v", err)
		return schedule, err
	}

	return schedule, nil

}

func GetScheduleByUser(u_id string) ([]*defs.Schedule, error) {
	scheduleList := make([]*defs.Schedule, 0)
	stmtOuts, err := dbConn.Prepare("SELECT `id`, `user_id`, `topic`, `content`,`created_at`, `status`, `active` FROM schedule WHERE `user_id` = ? AND `active` = 1 ORDER BY `created_at` DESC")
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
		err = rows.Scan(&schedule.Id, &schedule.UserId, &schedule.Topic, &schedule.Content, &t, &schedule.Status, &schedule.Active)
		schedule.CreatedAt = t.Format("2006-01-02 15:04:05")
		if err != nil {
			log.Printf("error: %+v", err)
			continue
		}
		schedule.UserId = u_id
		scheduleList = append(scheduleList, schedule)
	}

	defer stmtOuts.Close()

	return scheduleList, nil
}

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

func UpdateScheduleById(schedule *defs.Schedule) error {
	stmtIns, err := dbConn.Prepare("UPDATE schedule SET `active` = 0 WHERE `id` = ?")
	if err != nil {
		log.Printf("update schedule dbConn prepare error: %s\n", err)
		return err
	}

	_, err = stmtIns.Exec(schedule.Id)
	if err != nil {
		log.Printf("update schedule stmtIns exec error: %s\n", err)
		return err
	}

	stmtIns, err = dbConn.Prepare("INSERT INTO schedule (`user_id`, `topic`, `content`, `created_at`, `status`) VALUES(?, ?, ?, ?, ?)")
	if err != nil {
		log.Printf("add schedule dbConn prepare error: %s\n", err)
		return err
	}

	_, err = stmtIns.Exec(schedule.UserId, schedule.Topic, schedule.Content, schedule.CreatedAt, schedule.Status)
	if err != nil {
		log.Printf("add schedule stmtIns exec error: %s\n", err)
		return err
	}

	defer stmtIns.Close()

	return nil
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
