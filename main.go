package main

import (
	"fmt"
	"level-forum/fields"
	"net/http"
)

func main() {
	// 定义处理函数
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Hello, World!")
	})

	// 启动HTTP服务器
	port := 8080
	fmt.Printf("Server is listening on port %d...\n", port)
	err := http.ListenAndServe(fmt.Sprintf(":%d", port), nil)
	if err != nil {
		fmt.Println("Error:", err)
	}

	println("hello world")
	fields.CreateField("", nil)
}
