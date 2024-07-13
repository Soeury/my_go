package main

import (
	"fmt"
	"net/http"
)

func hello(writer http.ResponseWriter, request *http.Request) {

	fmt.Println(request.Method)
	fmt.Println(request.URL)
	fmt.Println(request.RemoteAddr)
	fmt.Println(request.Body)
	fmt.Println(request.Header)

	fmt.Println("Moonlight")

	writer.Write([]byte("I LOVE YOU FOREVER !"))
}

func main() {

	// 请求处理
	http.HandleFunc("/hello", hello)

	// 端口
	http.ListenAndServe("127.0.0.1:8080", nil)
}
