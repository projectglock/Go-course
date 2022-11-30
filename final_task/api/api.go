package api

import (
	"errors"
	"final_task/service"
	"final_task/tasksHandler"
	"fmt"
	"net/http"
	"net/url"
	"time"
)

// MainHandler - Cтруктура для обработчика точки входа
type MainHandler struct {
}

func (c MainHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	//декодируем кириллицу
	uriDecoded, err := url.PathUnescape(r.RequestURI)
	if err != nil {
		service.LogInternalError(w, errors.New("internal decoding error: "+err.Error()))
		return
	}

	//логгируем в консоль
	fmt.Println("New request", time.Now(), r.Method, uriDecoded)

	if r.Method == http.MethodGet {
		switch uriDecoded {
		case "/task/Циклическая ротация":
			tasksHandler.TasksHandler(w, "Циклическая ротация")
		case "/task/Чудные вхождения в массив":
			tasksHandler.TasksHandler(w, "Чудные вхождения в массив")
		case "/task/Проверка последовательности":
			tasksHandler.TasksHandler(w, "Проверка последовательности")
		case "/task/Поиск отсутствующего элемента":
			tasksHandler.TasksHandler(w, "Поиск отсутствующего элемента")
		}
	}

	if r.RequestURI == "/tasks" && r.Method == http.MethodGet {
		tasksHandler.TasksHandler(w, "Циклическая ротация")
		tasksHandler.TasksHandler(w, "Чудные вхождения в массив")
		tasksHandler.TasksHandler(w, "Проверка последовательности")
		tasksHandler.TasksHandler(w, "Поиск отсутствующего элемента")
		return
	}

	w.WriteHeader(http.StatusNotFound)
	return
}
