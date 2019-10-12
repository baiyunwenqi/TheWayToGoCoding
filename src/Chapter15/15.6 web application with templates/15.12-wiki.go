package main

import (
	"html/template"
	"io/ioutil"
	"net/http"
	"regexp"
)

const lenPath=len("/view/")

type Page struct {
	Title string
	Body []byte
}

func (p *Page) save() error{
	filename:=p.Title+".txt"
	return ioutil.WriteFile(filename,p.Body,0600) // read-write permissions
}
func load(title string) (*Page,error){
	filename:=title+".txt"
	body,err:=ioutil.ReadFile(filename)
	if err!=nil{
		return nil,err
	}
	return &Page{Title:title,Body:body},nil
}

var (
	titleValidator=regexp.MustCompile(`^[a-zA-Z0-9]+$`)
	templates = make(map[string]*template.Template)
	err error
)

func init(){
	for _,tmp:= range []string{"edit","view"}{
		templates[tmp]=template.Must(template.ParseFiles(tmp+".html"))
	}
}

func main(){
	http.HandleFunc("/view/",makeHandler(viewHandler))
	http.HandleFunc("/edit/",makeHandler(editHandler))
	http.HandleFunc("/save/",makeHandler(saveHandler))
}

func makeHandler(fn func(http.ResponseWriter,*http.Request,string)) http.HandlerFunc{
	return func(w http.ResponseWriter, r *http.Request){
		title:=r.URL.Path[lenPath:]
		if !titleValidator.MatchString(title){
			http.NotFound(w,r)
			return
		}
		fn(w,r,title)
	}

}

func viewHandler(w http.ResponseWriter, r *http.Request, title string){
	p,err:=load(title)
	if err!=nil{ // page not found redirect to edit page
		http.Redirect(w,r,"/edit/"+title,http.StatusFound)
		return
	}
	renderTemplate(w,"view",p) // 利用模板根据txt文件创建新的页面
}

func editHandler(w http.ResponseWriter, r *http.Request,title string){
	p,err:=load(title)
	if err!=nil{ // if not found, use an empty page
		p=&Page{Title:title}
	}
	renderTemplate(w,"edit",p)
}

func saveHandler(w http.ResponseWriter, r*http.Request, title string ){
	body:=r.FormValue("body")
	p:=&Page{Title:title,Body:[]byte(body)}
	err:=p.save()
	if err!=nil{
		http.Error(w,err.Error(),http.StatusInternalServerError)
		return
	}
	http.Redirect(w,r,"/view/"+title,http.StatusFound) // if cannot save successfully, redirect to the previous view page
}

func renderTemplate(w http.ResponseWriter, tmpl string, p *Page){
	err:=templates[tmpl].Execute(w,p)
	if err!=nil{
		http.Error(w, err.Error(),http.StatusInternalServerError)
	}

}