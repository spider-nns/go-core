package main

import (
	"html/template"
	"os"
	"strings"
)

func main() {
	funcMap := template.FuncMap{"title": strings.Title}
	tpl, _ := template.New("go").Funcs(funcMap).Parse(templateText)
	data := map[string]string{
		"Name1": "go",
		"Name2": "2021",
		"Name3": "初二",
	}
	tpl.Execute(os.Stdout, data)
}

const templateText = `
output 0:{{title.Name1}}
output 1:{{title.Name2}}
output 2:{{.Name3 | title}}
`
