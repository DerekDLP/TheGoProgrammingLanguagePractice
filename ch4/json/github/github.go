// 通过Github的issue查询服务
// 许多web服务都提供JSON接口，通过HTTP接口发送JSON格式请求并返回JSON格式的信息
//
// GitHub提供的Go API去追踪issue
// 参见：https://developer.github.com/v3/search/#search-issues

package github

import (
	"encoding/json"
	"net/http"
	"net/url"
	"strings"
	"time"
)

// IssuesURL 查询网址
const IssuesURL = "https://api.github.com/search/issues"

// IssuesSearchResult 查询结果
// 即使对应的JSON对象名是小写字母，每个结构体的成员名也是声明为大小字母开头的
type IssuesSearchResult struct {
	TotalCount int `json:"total_count"`
	Items      []*Issue
}

// Issue 格式
type Issue struct {
	Number    int
	HTMLURL   string `json:"html_url"`
	Title     string
	State     string
	User      *User
	CreatedAt time.Time `json:"created_at"`
	Body      string    //在Markdown格式
}

// User 格式
type User struct {
	Login   string
	HTMLURL string `json:"html_url"`
}

// SearchIssues 函数发出一个HTTP请求，然后解码返回的JSON格式的结果
func SearchIssues(terms []string) (*IssuesSearchResult, error) {
	// 因为用户提供的查询条件可能包含类似?和&之类的特殊字符
	// 为了避免对URL造成冲突，使用url.QueryEscape来对查询中的特殊字符进行转义操作
	q := url.QueryEscape(strings.Join(terms, " "))
	resp, err := http.Get(IssuesURL + "?q=" + q)
	if err != nil {
		return nil, err
	}

	// 必须关闭所有的执行路径上的resp.Body
	// 该需求可使用defer简化实现
	if resp.StatusCode != http.StatusOK {
		resp.Body.Close()
		return nil, err
	}

	var result IssuesSearchResult
	// 基于流式的解码器json.Decoder，它可以从一个输入流解码JSON数据
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		resp.Body.Close()
		return nil, err
	}
	resp.Body.Close()
	return &result, nil
}
