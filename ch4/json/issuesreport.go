// "text/template"模板
// go run issuesreport.go repo:golang/go is:open json decoder

package main

import (
	"log"
	"text/template"
	"time"

	"GoBook/TheGoProgrammingLanguagePractice/ch4/json/github"
	"os"
)

const templ = `{{.TotalCount}} issues:
{{range .Items}}----------------------------
Number: {{.Number}}
User:   {{.User.Login}}
Title:  {{.Title | printf "%.64s"}}
Age:    {{.CreatedAt | daysAgo}} days
{{end}}`

func daysAgo(t time.Time) int {
	return int(time.Since(t).Hours() / 24)
}

func main() {
	report, err := template.New("report").
		Funcs(template.FuncMap{"daysAgo": daysAgo}).
		Parse(templ)
	if err != nil {
		log.Fatal(err)
	}

	// 利用template.Must简化上述代码
	// var report = template.Must(template.New("issuelist").
	//     Funcs(template.FuncMap{"daysAgo": daysAgo}).
	// 	Parse(templ))

	result, err := github.SearchIssues(os.Args[1:])
	if err != nil {
		log.Fatal(err)
	}
	if err := report.Execute(os.Stdout, result); err != nil {
		log.Fatal(err)
	}
}
