package main

import (
	"fmt"
	"net/http"
	"runtime"
)

func main() {
	// 设置路由规则
	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Fprintf(writer, "你好，世界！"+runtime.GOOS+" "+runtime.GOARCH)
	})
	// 启动HTTP服务
	http.ListenAndServe(":6900", nil)
}
