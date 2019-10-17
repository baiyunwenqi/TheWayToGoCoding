package main

import(
	"log"
	"net/smtp"	
	"bytes"
)

func main(){
	client, err := smtp.Dial("mail.example.com:25")
	if err!=nil{
		log.Fatal(err)
	}
	// Set the sender and recipient.
	client.Mail("982806073@qq.com0")
	client.Rcpt("liwenqi32@jd.com")
	// send the email body
	wc,err:=client.Data()
	if err!=nil{
		log.Fatal(err)
	}
	defer wc.Close()
	buf:=bytes.NewBufferString("this is the email body.")
	if _,err=buf.WriteTo(wc);err!=nil{
		log.Print("error")
	}
}

