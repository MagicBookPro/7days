package main

import (
	"fmt"
	"log"
	"net/http"
)

//在实现Engine之前，我们调用 http.HandleFunc 实现了路由和Handler的映射，也就是只能针对具体的路由写处理逻辑。比如/hello。
//但是在实现Engine之后，我们拦截了所有的HTTP请求，拥有了统一的控制入口。在这里我们可以自由定义路由映射的规则，
//也可以统一添加一些处理逻辑，例如日志、异常处理等。
type Engine struct {
}

func (e *Engine) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	switch request.URL.Path {
	case "/":
		fmt.Fprintf(writer, "URL.Path = %q\n", request.URL.Path)
	case "/hello":
		for k, v := range request.Header {
			fmt.Fprintf(writer, "Header[%q] =  %q\n", k, v)
		}
	default:
		fmt.Fprintf(writer, "404 not found %s\n", request.URL)
	}
}

func main() {
	engine := new(Engine)
	log.Fatal(http.ListenAndServe(":9999", engine))
}
