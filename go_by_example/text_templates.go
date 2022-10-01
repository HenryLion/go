package main

import (
	"os"
	"text/template"
)

func main() {
	t1 := template.New("t1shadongxi")
	temp, err := t1.Parse("Value is: {{.}}\n")
	if err != nil {
		panic(err)
	}

	slice1 := []string{"Go", "Rust", "C++", "C#"}

	temp.Execute(os.Stdout, "some text")
	temp.Execute(os.Stdout, 5)
	temp.Execute(os.Stdout, slice1)

	t2 := template.New("template2")
	temp2 := template.Must(t2.Parse("Value2: {{.}}\n"))
	temp2.Execute(os.Stdout, slice1)

	Create := func(name, t string) *template.Template {
		return template.Must(template.New(name).Parse(t))
	}

	t3 := Create("template3", "Name: {{.Name}}\n")

	t3.Execute(os.Stdout, struct {
		Name string
	}{"Jane Doe"})

	t3.Execute(os.Stdout, map[string]string{
		"Name": "Mickey Mouse",
	})

	t4 := Create("template4",
		"{{if . -}} yes {{else -}} no {{end}}\n")

	t4.Execute(os.Stdout, "not empty")
	t4.Execute(os.Stdout, "")
}
