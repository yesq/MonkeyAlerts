package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"

	"time"

	gomail "gopkg.in/gomail.v2"
)

// Config : smtp config
type Config struct {
	Mail     string `json:"mail"`
	Password string `json:"password"`
	SMTPURL  string `json:"smtpURL"`
	SMTPPort int    `json:"smtpPort"`
}

// CH : channel to transport mail
var CH chan *gomail.Message

var config Config

var d *gomail.Dialer

func init() {
	loadConfig()
	fmt.Println(config)
	go daemonMailClient()
}

func loadConfig() error {
	fl, err := os.Open("./config.json")
	if err == nil {
		str, err := ioutil.ReadAll(fl)
		if err != nil {
			return fmt.Errorf("read ./config.json err : %v", err)
		}
		return json.Unmarshal([]byte(str), &config)
	}
	defer fl.Close()
	return fmt.Errorf("open ./config.json err : %v", err)
}

func sendAlertSample(addTo, body, title string) {
	println("send line 1")
	m := gomail.NewMessage()
	m.SetHeader("From", config.Mail)
	m.SetHeader("To", addTo)
	m.SetHeader("Subject", title)
	m.SetBody("text/html", body)
	// CH <- m
	println("send CH")
	CH <- m
}

func daemonMailClient() {
	println("daemon start")
	CH = make(chan *gomail.Message)
	defer close(CH)
	defer println("exec defer")
	d = gomail.NewDialer(config.SMTPURL, config.SMTPPort, config.Mail, config.Password)
	var s gomail.SendCloser
	var err error
	open := false
	println("for start")
	for {
		println("for 1")
		select {
		case m, ok := <-CH:
			println("<-CH")
			if !ok { // 是否 close(CH)
				return
			}
			if !open {
				if s, err = d.Dial(); err != nil {
					panic(err)
				}
				open = true
			}
			if err := gomail.Send(s, m); err != nil {
				println(err)
				// log.Print(err) //log?
			}
		case <-time.After(30 * time.Second):
			println("<-timeout")
			if open {
				if err := s.Close(); err != nil {
					panic(err)
				}
				open = false
			}
		}
	}
}

func main() {
	for {
		select {
		case <-time.After(40 * time.Second):
			sendAlertSample("1508866205@qq.com", "test", "title")
		}
	}
}
