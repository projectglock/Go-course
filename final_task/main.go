package main

import (
	"final_task/service"
	"net/http"
)

func main() {
	err := http.ListenAndServe("localhost:3000", service.MyHandler{})
	if err != nil {
		panic(err)
	}
}
