package main

import (
	"fmt"
	"io"
	"net/http"
)

func main() {

	resp, err := http.Get("http://127.0.0.1:8080/hello")
	if err != nil {
		fmt.Println("found err :", err)
	}
	defer resp.Body.Close()

	buf := make([]byte, 1024)

	for {
		n, err := resp.Body.Read(buf)
		if err != nil && err != io.EOF {
			fmt.Println(err)
			return
		} else {
			fmt.Println("fetch over!")
			fmt.Println(string(buf[:n]))
		}
	}
}
