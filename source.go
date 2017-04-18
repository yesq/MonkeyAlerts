package main

import (
	"database/sql"
	"strconv"
	"time"
)

// TODO : change to Redis

var db *sql.DB

func DB() *sql.DB {
	if db != nil {
		return db
	} else {
		var err error
		db, err = sql.Open("mysql", config.MySQL)
		// db.SetMaxOpenConns(10)
		checkErr(err)
		return db
	}
}

// GetTarget : source - target
func GetSourceTarget(source string) (target string, touchLimit bool, ok bool) {
	db := DB()
	stmt, err := db.Prepare("UPDATE `source` SET `count` = `count` + 1 WHERE `source` = ?")
	checkErr(err)
	res, err := stmt.Exec(source)
	checkErr(err)
	num, _ := res.RowsAffected()
	if num == 1 {
		target, touchLimit := GetSourceInfo(source)
		return target, touchLimit, true
	}
	return "", false, false
}

// GetSourceInfo
// target
// whether mail available
func GetSourceInfo(source string) (string, bool) {
	db := DB()
	rows, err := db.Query("SELECT `target`, `count`, `countLimit`, `intervalLimit`, `lastAlert` FROM `source` WHERE `source` = \"" + source + "\";")
	var touchLimit bool
	var target string
	for rows.Next() {
		now := int(time.Now().Unix())
		var count int
		var countLimit int
		var intervalLimit int
		var lastAlert int
		err = rows.Scan(&target, &count, &countLimit, &intervalLimit, &lastAlert)
		checkErr(err)
		touchCountsLimit := count > countLimit
		touchTimeLimit := now > (lastAlert + intervalLimit)
		if touchCountsLimit && touchTimeLimit {
			resetCount(source)
			touchLimit = true
		}
	}
	return target, touchLimit
}

func resetCount(source string) {
	db := DB()
	now := int(time.Now().Unix())
	_, err := db.Query("UPDATE `source` SET `count`=0, `lastAlert`=" + strconv.Itoa(now) + " WHERE `source`=\"" + source + "\"")
	checkErr(err)
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
