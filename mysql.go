package main

import (
	"database/sql"
	"errors"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

// GetMysqlConn gets a mysql connection
func GetMysqlConn() (*sql.DB, error){
	db, err := sql.Open("mysql", Conf.Mysql.Dsn)

	if err != nil {
		fmt.Println("get mysql conn error, ", err)
		return db, err
	}

	return db, nil
}

// SaveRecords saves winning records
func SaveRecords(awardName string, awardTime string, userName string) error {
	db, err := GetMysqlConn()
	if err != nil {
		fmt.Println("conn is nil")
		return err
	}

	defer db.Close()

	stmt, err := db.Prepare("insert into award_user_info(award_name,award_time,user_name) values(?,?,?);")
	if err != nil {
		fmt.Println("prepare insert sql error, ", err)
		return errors.New("prepare insert sql error")
	}

	fmt.Println("insert into award_user_info , award_name , award_time  , user_name ",awardName, awardTime, userName)

	_ , err = stmt.Exec(awardName, awardTime, userName)
	if err != nil {
		fmt.Println("exec sql error, ", err)
		return errors.New("exec sql error")
	}

	return nil
}

// QueryRecords querys winning records
func QueryRecords() error {

	db, err := GetMysqlConn()
	if err != nil {
		fmt.Println("conn is nil")
		return err
	}

	defer db.Close()

	rows, err := db.Query("select * from award_user_info")
	if err != nil {
		fmt.Println("exec select error", err)
		return err
	}

	for rows.Next() {
		var id int64
		var awardName string
		var userName string
		var awardTime string

		err = rows.Scan(&id, &awardName, &userName, &awardTime)
		fmt.Printf("id : %d, awardName : %s, userName : %s, awardTime : %s \n",id, awardName, userName, awardTime)
	}

	return err
}

