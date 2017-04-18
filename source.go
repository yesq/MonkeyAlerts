package main

import "database/sql"

// TODO : change to Redis

var db *sql.DB

func DB() *sql.DB {
	if db != nil {
		return db
	} else {
		var err error
		db, err = sql.Open("mysql", config.MySql)
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
// whether touch limit
func GetSourceInfo(source string) (string, bool) {
	db := DB()
	rows, err := db.Query("SELECT `target`, `count`, `countLimit` FROM `source` WHERE `source` = \"" + source + "\";")
	var touchLimit bool
	var target string
	var count int
	var countLimit int
	for rows.Next() {
		err = rows.Scan(&target, &count, &countLimit)
		checkErr(err)
		if count > countLimit {
			resetCount(source)
			touchLimit = true
		}
	}
	return target, touchLimit
}

func resetCount(source string) {
	db := DB()
	_, err := db.Query("UPDATE `source` SET `count`=0 WHERE `source`=\"" + source + "\"")
	checkErr(err)
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
