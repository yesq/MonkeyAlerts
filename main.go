package main

import gomail "gopkg.in/gomail.v2"
import "crypto/tls"

func main() {
	d := gomail.NewDialer("smtp.qq.com", 465, "username", "password")
	d.TLSConfig = &tls.Config{InsecureSkipVerify: true}

	m := gomail.NewMessage()
	m.SetHeader("From", "****@qq.com")
	m.SetHeader("To", "****@qq.com", "****@gmail.com")
	m.SetBody("text/html", "Hello <b>YesQ</b>")

	if err := d.DialAndSend(m); err != nil {
		panic(err)
	}
}
