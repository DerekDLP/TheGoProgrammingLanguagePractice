// echo HTTP请求
package main
import (
	"fmt"
	"log"
	"net/http"
)

func main()  {
	// 每个请求的处理句柄
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

// echo HTTP请求
func handler(w http.ResponseWriter, r *http.Request)  {
	fmt.Fprintf(w, "%s  %s  %s\n", r.Method, r.URL, r.Proto)
	for k, v := range r.Header {
		fmt.Fprintf(w, "Header[%q] = %q\n", k, v)
	}
	fmt.Fprintf(w, "Host = %q\n", r.Host)
	fmt.Fprintf(w, "RemoteAddr = %q\n", r.RemoteAddr)
	// 与下述if等价
	// err := r.ParseForm()
	// if rr != nil {
	//     log.Print(err) 
	// }
	if err := r.ParseForm(); err != nil {
		log.Print(err) 
	}
	for k, v := range r.Form {
		fmt.Fprintf(w, "Form[%q] = %q\n", k, v)
	}
}