package main

import (
	"fmt"
	"net/http"
)

const DefaultPort = ":8080"


// リクエストを処理する関数
func handler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "test page from Go.")
}


func main() {

    // パスに対して処理を追加
    http.HandleFunc("/", handler)

    // 8080ポートで起動
    http.ListenAndServe(":8080", nil)
}