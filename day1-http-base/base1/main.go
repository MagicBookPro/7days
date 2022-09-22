package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", indexHandler)
	//http.HandleFunc 实现了路由和Handler的映射，也就是只能针对具体的路由写处理逻辑。比如/hello。
	http.HandleFunc("/hello", helloHandler)
	//第一个参数是地址，:9999表示在 9999 端口监听。而第二个参数则代表处理所有的HTTP请求的实例，nil 代表使用标准库中的实例处理。
	log.Fatal(http.ListenAndServe(":9999", nil))

}

func helloHandler(writer http.ResponseWriter, request *http.Request) {

	for k, v := range request.Header {
		fmt.Fprintf(writer, "Head[%q] = %q\n", k, v)
	}
}

func indexHandler(writer http.ResponseWriter, request *http.Request) {
	//fmt.Fprintf() 依据指定的格式向第一个参数内写入字符串，第一个参数必须实现了 io.Writer 接口。
	//%q 带双引号的字符串"abc"或带单引号的字符'c'
	//fmt.Fprintf(writer, "URL = %q\n", request.URL)
	fmt.Fprintf(writer, "URL.Path = %q\n", request.URL.Path)
	//fmt.Fprintf(writer, "URL.Host = %q\n", request.URL.Host)
	//fmt.Fprintf(writer, "URL.RawPath = %q\n", request.URL.RawPath)
	//fmt.Fprintf(writer, "Header = %q\n", request.Header)
	//fmt.Fprintf(writer, "Host = %q\n", request.Host)
	//fmt.Fprintf(writer, "Body = %q\n", request.Body)
	//fmt.Fprintf(writer, "Method = %q\n", request.Method)
}
