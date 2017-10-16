package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

// ルーティング
type Routing struct {
	Routes []Route //`json:routes`
}

// ルート
type Route struct {
	Url      string //`json:url`
	Response string //`json:response_json`
}

// バージョン
func version() {
	fmt.Printf("Version: v1.0.0\n")
}

// json読み込み
func readJson() Routing {
	fmt.Printf("Reading json...\n")

	// 設定ファイルを読み込む
	file, e := ioutil.ReadFile("./config.json")
	if e != nil {
		fmt.Printf("File error: %v\n", e)
		os.Exit(1)
	}

	// パース
	var routing Routing
	json.Unmarshal(file, &routing)
	return routing

}

// サーバー起動
func startServe(routing Routing) {
	fmt.Printf("Starting server...\n")

	// ルーティング設定
	fmt.Printf("Route >>> \n")
	for _, route := range routing.Routes {
		fmt.Printf("   %v\n", route.Url)
		response := route.Response
		http.HandleFunc(route.Url, func(w http.ResponseWriter, req *http.Request) {
			w.Write([]byte(response))
		})
	}

	// サーバー起動
	http.ListenAndServe(":8080", nil)
	fmt.Printf("Started server => http://localhost:8080\n")

}

func main() {
	version()
	fmt.Printf("Starting quickry...\n")

	// ルーティング設定ファイルの内容を取得
	routing := readJson()

	// サーバー起動
	startServe(routing)

}
