package main

import (
    "fmt"
    "net/http"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
    name := r.URL.Query().Get("name") // 获取查询参数 "name"
    if name == "" {
        name = "World"
    }
    fmt.Fprintf(w, "Hello, %s!", name)
}

func main() {
    http.HandleFunc("/", helloHandler) // 设置根路径的路由
    fmt.Println("服务器正在 http://localhost:8080 运行...")
    http.ListenAndServe(":8080", nil)
}
