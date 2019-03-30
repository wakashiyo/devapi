package main

import (
	"fmt"
	"net/http"
)

//handler : http request handler, return "hello world"
func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, World !!!!!!!!!!!!!!!!!!!")
}

func main() {
	http.HandleFunc("/", handler) // ハンドラを登録してウェブページを表示させる
	http.ListenAndServe(":9001", nil)
}
