package http

import (
	"net/http"
)

func (h*HttpServer)run(op func()) {
	//http://127.0.0.1:8000/go
	// 单独写回调函数
	http.HandleFunc("/go", Handler)
	//http.HandleFunc("/ungo",myHandler2 )
	// addr：监听的地址
	// handler：回调函数
	http.ListenAndServe("127.0.0.1:8000", nil)
}

