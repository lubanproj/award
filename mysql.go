package main

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

// GetMysqlConn gets a mysql connection
func GetMysqlConn() *sql.DB{
	db, err := sql.Open("mysql", Conf.Mysql.Dsn)

	if err != nil {
		log.Println("get mysql conn error, ", err)
		return nil
	}

	return db
}

// SaveRecords keeps winning records
func SaveRecords(awardName string, awardTime string, userName string) {
	db := GetMysqlConn()

	defer db.Close()

	if db == nil {
		log.Println("conn is nil")
		return
	}

	stmt, err := db.Prepare("insert into award_user_info(award_name,award_time,user_name) values(?,?,?);")
	if err != nil {
		log.Println("prepare insert sql error, ", err)
		return
	}

	log.Println("insert into award_user_info , award_name , award_time  , user_name ",awardName, awardTime, userName)

	_ , err = stmt.Exec(awardName, awardTime, userName)
	if err != nil {
		log.Println("exec sql error, ", err)
		return
	}

}

// QueryRecords querys winning records
func QueryRecords() {

	db := GetMysqlConn()
	defer db.Close()

	if db == nil {
		log.Println("conn is nil")
		return
	}

	rows, err := db.Query("select * from award_user_info")
	if err != nil {
		log.Println("exec select error", err)
		return
	}

	for rows.Next() {
		var id int64
		var awardName string
		var userName string
		var awardTime string

		err = rows.Scan(&id, &awardName, &userName, &awardTime)
		log.Printf("id : %d, awardName : %s, userName : %s, awardTime : %s \n",id, awardName, userName, awardTime)
	}

}

