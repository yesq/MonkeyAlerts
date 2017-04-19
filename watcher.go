package main

import (
	"bytes"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
	"time"
)

type requestTask struct {
	Method      string
	URL         string
	Body        string
	Headers     string
	Timeout     int
	RightBody   string
	RightStatus int
	Source      string
}

func daemonWatcher() {
	for {
		select {
		case <-time.After(3 * time.Second):
			loadTasks()
		}
	}
}

func loadTasks() {
	db := DB()
	rows, err := db.Query("select r.method, r.URL, r.body, r.headers, r.timeout, r.rightStatus, r.rightBody, s.source from `requestTast` r inner join `source` s;")
	checkErr(err)
	defer rows.Close()
	for rows.Next() {
		var r requestTask
		rows.Scan(&r.Method, &r.URL, &r.Body, &r.Headers, &r.Timeout, &r.RightStatus, &r.RightBody, &r.Source)
		execTask(r)
	}
}

func setHeaders(r requestTask, req *http.Request) {
	if r.Headers != "" {
		rows := strings.Split(r.Headers, ";")
		for _, v := range rows {
			head := strings.Split(v, ":")
			req.Header.Set(head[0], head[1])
		}
	}
}

func watcherAlert(source string, errBody string) {
	if target, touchLimit, ok := GetSourceTarget(source); ok {
		if touchLimit {
			println("send mail")
			sendAlertSample(target, errBody, "service error from "+source)
			resetCount(source)
		}
	}
}

func execTask(r requestTask) error {
	req, err := http.NewRequest(r.Method, r.URL, bytes.NewBuffer([]byte(r.Body)))
	checkErr(err)
	setHeaders(r, req)
	body, status, err := doRequest(req)
	if err != nil {
		watcherAlert(r.Source, err.Error())
	}
	if status != r.RightStatus || (r.RightBody != "" && (string(body) != r.RightBody)) {
		println("alert")
		watcherAlert(r.Source, strconv.Itoa(status)+":"+string(body))
	}
	return err
}

func doRequest(req *http.Request) ([]byte, int, error) {
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, 0, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	checkErr(err)
	return body, resp.StatusCode, nil
}
