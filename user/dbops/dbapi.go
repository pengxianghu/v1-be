package dbops

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"log"

	"github.com/pengxianghu/v1-be/user/defs"
)

func AddUser(user *defs.User) error {
	stmtIns, err := dbConn.Prepare("INSERT IGNORE INTO user (`id`, `name`, `pwd`) VALUES(?, ?, ?)")
	if err != nil {
		log.Printf("add user dbConn prepare error: %s\n", err)
		return err
	}

	_, err = stmtIns.Exec(user.Id, user.Name, user.Pwd)
	if err != nil {
		log.Printf("add user stmtIns exec error: %s\n", err)
		return err
	}

	defer stmtIns.Close()

	return nil
}

func GetUserCredential(u_name string) (*defs.User, error) {
	stmtOuts, err := dbConn.Prepare("SELECT `id`,`name`,`pwd` FROM user WHERE `name` = ?")
	if err != nil {
		log.Printf("get user stmtOuts error: %s\n", err)
		return &defs.User{}, err
	}

	db_user := &defs.User{}
	err = stmtOuts.QueryRow(u_name).Scan(&db_user.Id, &db_user.Name, &db_user.Pwd)
	// 不为空且不是no rows
	if err != nil && err != sql.ErrNoRows {
		return &defs.User{}, err
	}
	defer stmtOuts.Close()

	return db_user, nil
}
