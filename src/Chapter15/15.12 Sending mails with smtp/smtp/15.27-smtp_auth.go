package main

import(
	"log"
	"net/smtp"
)

func main(){
	auth:=smtp.PlainAuth(
		"",
		"982806073@qq.com",
		"practice123,.",
		"mail.qq.com",
	)
	// connect to server, authenticate, set the sender and recipient and send the email all in one step
	err:= smtp.SendMail(
		 "982806073@qq,com",
		auth,
		"liwenqi32@jd.com",
		[]string{"recipient@example.net"},
		[]byte("This is the email body."),
	)
	if err!=nil{
		log.Print(err)
	}
}