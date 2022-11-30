package main

import (
	"final_task/api"
	"net/http"
)

func main() {
	err := http.ListenAndServe("localhost:3000", api.MainHandler{})
	if err != nil {
		panic(err)
	}
}
