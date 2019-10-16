package main

import (
	"log"
	"net/http"
	"fmt"
	"strings"
	"strconv"
	"sort"
)

type statistics struct{
	numbers []float64
	mean float64
	median float64
}

const body= `  <head><meta http-equiv="content-type" content="text/html; charset=UTF-8"></head>
<h1>Statistics</h1><p>Computes basic statistics for a given list of numbers<br/></p>`

const form= 	`<html>
<body>
<form id="myform" action="/" method="POST" target="iframe">
<label for="number">Numbers(comma or space-separated):</label><br>
<input type="text name="numbers" size="30"><br />
<input type="submit" value="Calculate">
</form>
<iframe id="iframe" name="iframe" style="display:none;"></iframe>
</body>
</html>`

const error = `<p class="error">%s</p>`

var pageTop=body
var pageBottom=""

func ComputingHandle(w http.ResponseWriter, request *http.Request){
	w.Header().Set("Content-Type","text/html")
	switch request.Method{
	case "GET":
		fmt.Fprint(w,pageTop,form)
		log.Print("A GET request")
		err:=request.ParseForm()
		if err!=nil{
			fmt.Fprintf(w,error,err)
			log.Printf(err.Error())
		}else{
			if numbers,message,ok:=ParseRequest(request);ok{
				stats:=getStats(numbers)
				fmt.Fprint(w,formatStats(stats))
			}else if message !=""{
				fmt.Fprintf(w,error,message)
			}
		}
	case "POST":
		log.Print("A post request")
		fmt.Fprint(w,pageTop,form)
		err:=request.ParseForm()
		if err!=nil{
			fmt.Fprintf(w,error,err)
			log.Printf(err.Error())
		}else{
			if numbers,message,ok:=ParseRequest(request);ok{
				stats:=getStats(numbers)
				fmt.Fprint(w,formatStats(stats))
			}else if message !=""{
				fmt.Fprintf(w,error,message)
			}
		}
	}
	fmt.Fprint(w,pageBottom)
	return
}

func ParseRequest(request *http.Request) (numbers []float64, errorMessage string, isEmpty bool){
	if slice,found:=request.Form["numbers"]; found && len(slice)>0 {
		text := strings.Replace(slice[0],","," ",-1)
		for _,field := range strings.Fields(text) {
			if x ,err := strconv.ParseFloat(field,64);err!=nil{
				return numbers,"'" + field +"'" + "is invalid", false
			}else{
				numbers=append(numbers,x)
			}
		}
	}
	if len(numbers)==0{
		return numbers, "",false
	}
	return numbers,"",true
}

func getStats(numbers []float64) (stat statistics){
	stat.numbers=numbers
	sort.Float64s(stat.numbers)
	stat.mean= sum(numbers) / float64(len(numbers))
	stat.median= median(stat.numbers)
	return 
}

func sum(numbers []float64) (res float64){
	for _,x:= range numbers{
		res+=x
	}
	return 
}

func median(numbers []float64) (res float64){
	mid:=len(numbers)/2
	res=numbers[mid]
	if len(numbers)%2==0{
		res=(res+numbers[mid-1])/2
	}
	return 
}

func formatStats(stat statistics) string{
	return fmt.Sprintf(`<table border="1">
<tr><th colspan="2">Results</th></tr>
<tr><td>Numbers</td><td>%v</td><tr>
<tr><td>Count</td><td>%d</td></tr>
<tr><td>Mean</td><td>%f</td></tr>
<tr><td>Median</td><td>%f</td></tr>
</table>`,stat.numbers,len(stat.numbers),stat.mean,stat.median)
}

func main(){
	http.HandleFunc("/",ComputingHandle)
	if err:=http.ListenAndServe(":38088",nil);err!=nil{
		log.Fatalf("unable to listen and server at port 38088")
	}
}