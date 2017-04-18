package main

import (
	"crypto/tls"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"

	gomail "gopkg.in/gomail.v2"
)

// Config : smtp config
type Config struct {
	Mail     string `json:"mail"`
	Password string `json:"password"`
	SMTPURL  string `json:"smtpURL"`
	SMTPPort int    `json:"smtpPort"`
}

var config Config
var d *gomail.Dialer

func init() {
	loadConfig()
	fmt.Println(config)
	d = gomail.NewDialer(config.SMTPURL, config.SMTPPort, config.Mail, config.Password)
	d.TLSConfig = &tls.Config{InsecureSkipVerify: true}
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

func sendAlertSample(addTo string, body string) error {
	m := gomail.NewMessage()
	m.SetHeader("From", config.Mail)
	m.SetHeader("To", addTo)
	m.SetBody("text/html", body)
	return d.DialAndSend(m)
}

func main() {
	sendAlertSample("1508866205@qq.com", "test")
}
