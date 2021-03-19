package main

import (
	"fmt"
	"net/smtp"
)

func main() {
	send("bill.rassel@gmail.com", "Bismillah testing from golang")
}

//before sending email
//make sure your  2step security is turned on/enabled
//and app password is generated
func send(to, body string) bool {

	from := "yourusername@gmail.com" //yourusername@gmail.com
	pass := "yourAppPassword"        //app specific password yourAppPassword
	//to := "bill.rassel@gmail.com" //receiver

	msg := "From: " + from + "\n" +
		"To: " + to + "\n" +
		"Subject: Hello mostain\n\n" +
		body

	err := smtp.SendMail("smtp.gmail.com:587",
		smtp.PlainAuth("", from, pass, "smtp.gmail.com"),
		from, []string{to}, []byte(msg))

	if err != nil {
		fmt.Println("SMTP error", err.Error())
		return false
	}
	fmt.Println("email successfully sent")
	return true
}
