// server2是一个迷你的回声和计数器服务器
// 此程序没有做到在访问 /count 时不进行计数
package main

import (
	"fmt"
	"log"
	"net/http"
	"sync"
)

var mu sync.Mutex
var count int

func main() {
	http.HandleFunc("/", handler)
	http.HandleFunc("/count", counter)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

// 回显请求的 URL 的路径部分
func handler(w http.ResponseWriter, r *http.Request) {
	mu.Lock()
	fmt.Fprintf(w, "%s %s %s\n", r.Method, r.URL, r.Proto)
	for k, v := range r.Header {
		fmt.Fprintf(w, "Header[%q] = %q\n", k, v)
	}
	fmt.Fprintf(w, "RemoteAddr = %q\n", r.Host)
	fmt.Fprintf(w, "RemoteAddr = %q\n", r.RemoteAddr)
	if err := r.ParseForm(); err != nil {
		log.Print(err)
	}
	for k, v := range r.Form {
		fmt.Fprintf(w, "Form[%q] = %q\n", k, v)
	}
	count++
	fmt.Println(r.URL.Path)
	mu.Unlock()
	fmt.Fprintf(w, "RUL.Path = %q\n", r.URL.Path)
}

// counter 回显目前为止调用的次数
func counter(w http.ResponseWriter, r *http.Request) {
	mu.Lock()
	fmt.Fprintf(w, "Count %d\n", count)
	mu.Unlock()
}
