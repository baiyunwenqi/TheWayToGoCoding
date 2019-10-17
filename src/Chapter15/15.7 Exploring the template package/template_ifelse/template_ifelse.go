package main
import(
	"os"
	"text/template"
)
func main(){
	tEmpty := template.New("template test")
	tEmpty = template.Must(tEmpty.Parse("Empty pipleline if demo: {{if ``}} Will not print.{{end}}\n")) // the condition is false so nothing will be print
	tEmpty.Execute(os.Stdout,nil)

	tWithValue:= template.New("template test")
	tWithValue = template.Must( tWithValue.Parse("Non empty pipline if demo:{{if `anything`}} Will Print.{{end}}\n"))
	tWithValue.Execute(os.Stdout,nil)

	tIfElse:= template.New("template test")
	tIfElse = template.Must(tIfElse.Parse("if-else demo:{{if `anything`}} Print IF part.{{else}} Print ELSE part.{{end}}\n"))
	tIfElse.Execute(os.Stdout,nil)
}