package main

import (
	"net/http"
)

func main() {
	err := http.ListenAndServe("localhost:3000", candy.NewService())
	if err != nil {
		panic(err)
	}
}
