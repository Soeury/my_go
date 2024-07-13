package main

import (
	"fmt"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.Method)
	fmt.Println("----1")
	fmt.Println(r.URL)
	fmt.Println("----2")
	fmt.Println(r.RemoteAddr)
	fmt.Println("----3")
	fmt.Println(r.Body)
	fmt.Println("----4")

	fmt.Fprint(w, "hello world!") // fmt.Fprint( io.writer , 参数..) 将文本输出到指定的 io.writer 接口
	w.Write([]byte("I LOVE YOU FOREVER !"))
}

func main() {

	sever := http.Server{
		Addr: "0.0.0.0:80",
	}

	http.HandleFunc("/", handler)
	sever.ListenAndServe()
}
