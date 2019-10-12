package main

import (
	"io"
	"log"
	"net/http"
)

const body= `<h1>Statistics</h1><p>Computes basic statistics for a given list of numbers<br/></p>`

const form= 	`<html><body><form action="#" method="post" name="bar">
				Numbers(comma or space-separated):<br>
				<input type="text name="numbers/>
				<br><br>
                <input type="submit" value="Calculate"/
				</form></html></body>`

func ComputingHandle(w http.ResponseWriter, request *http.Request){
	w.Header().Set("Content-Type","text/html")
	switch request.Method {
	case "GET":
		_, _ = io.WriteString(w, body)
		_, _ = io.WriteString(w, form)
	case "POST":
	}
	return
}
func main(){
	http.HandleFunc("/compute",ComputingHandle)
	if err:=http.ListenAndServe(":8088",nil);err!=nil{
		log.Fatalf("unable to listen and server at port 8088")
	}
}