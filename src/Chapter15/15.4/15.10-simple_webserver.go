package main

import (
	"io"
	"log"
	"net/http"
)

const form= 	`<html><body><form action="#" method="post" name="bar">
				<input type="text name="in/>
                <input type="submit" value="Submit"/
				</form></html></body>`

func SimpleServer(w http.ResponseWriter, request *http.Request){
	if _,err:=io.WriteString(w,"<h1>hello, world</h1>");err!=nil{
		log.Fatalf("error")
		return
	}
}

/* handle a form, both the get which displays the form and the post which processes it.
 */
func FormServer(w http.ResponseWriter, request *http.Request){
	w.Header().Set("Content-Type","text/html")
	switch request.Method{
	case "GET":
		_,_=io.WriteString(w,form)
	case "POST":
		/* handle the form data, note that ParseForm must
		be called before we can extract form data*/
		//request.ParseForm();
		//io.WriteString(w, request.Form[“in”][0])
		_,_=io.WriteString(w,request.FormValue("in"))
	}
}

func main(){
	http.HandleFunc("/test1",SimpleServer)
	http.HandleFunc("/test2",FormServer)
	if err:=http.ListenAndServe(":8088",nil);err!=nil{
		log.Fatalf("unable to listen and server at port 8088")
	}
}
