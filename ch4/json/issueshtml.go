// "html/template"模板
// 只会没有包含会对HTML格式产生冲突的特殊字符
// go run issueshtml.go commenter:gopherbot json encoder >issues.html
// 无法
// go run issueshtml.go 3133 10535 >issues2.html

package main

import (
	"html/template"
	"log"
	"time"

	"GoBook/TheGoProgrammingLanguagePractice/ch4/json/github"
	"os"
)

func daysAgo(t time.Time) int {
	return int(time.Since(t).Hours() / 24)
}

func main() {
	var issueList = template.Must(template.New("issuelist").Parse(`
	<h1>{{.TotalCount}} issues</h1>
	<table>
	<tr style='text-align: left'>
	  <th>#</th>
	  <th>State</th>
	  <th>User</th>
 	  <th>Title</th>
	</tr>
	{{range .Items}}
	<tr>
	<td><a href='{{.HTMLURL}}'>{{.Number}}</a></td>
	<td>{{.State}}</td>
	<td><a href='{{.User.HTMLURL}}'>{{.User.Login}}</a></td>
	<td><a href='{{.HTMLURL}}'>{{.Title}}</a></td>
	</tr>
	{{end}}
	</table>
	`))

	result, err := github.SearchIssues(os.Args[1:])
	if err != nil {
		log.Fatal(err)
	}
	if err := issueList.Execute(os.Stdout, result); err != nil {
		log.Fatal(err)
	}
}
