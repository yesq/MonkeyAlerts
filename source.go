package main

import "database/sql"

func DB() *sql.DB {
	db, err := sql.Open("mysql", config.MySql)
	checkErr(err)
	return db
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
			touchLimit = true
		}
	}
	return target, touchLimit
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
