package service

import (
	"net/http"
)

func LogInternalError(w http.ResponseWriter, err error) {
	if err == nil {
		return
	}

	w.WriteHeader(http.StatusInternalServerError)
	w.Write([]byte(err.Error()))
}
