package main

import (
	"os"
	"text/template"
	"log"
	"encoding/json"
)
type Person struct{
	Name string
	nonExportedAgeField string
}
func main(){
	log.Println("New hello template")
	t:=template.New("hello")
	sv,_:=json.Marshal(t)
	log.Printf("template hello after New: %s",string(sv))
	t,_ = t.Parse("hello {{.Name}}!")
	sv,_=json.Marshal(t)
	log.Printf("template hello after Parse: %s",string(sv))
	p:= Person{Name: "Mary", nonExportedAgeField: "31"}
		if err:= t.Execute(os.Stdout,p);err!=nil{
		log.Printf("Error: %s",err.Error())
	}
}