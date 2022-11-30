package service

import (
	"encoding/json"
	"errors"
	"final_task/internal"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

// MyHandler - Cтруктура для обработчика
type MyHandler struct {
}

type BodyToSolutions struct {
	UserName string  `json:"user_name"` //имя юзера указанное в тг
	Task     string  `json:"task"`      //"имя задачи",
	Results  Results `json:"results"`
}

type Results struct {
	Payload string  `json:"payload"` // данные полученные для решения задачи
	Results [][]int `json:"results"` // результаты полученные при решении задачи
}

func (c MyHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	//декодируем кириллицу
	uriDecoded, err := url.PathUnescape(r.RequestURI)
	if err != nil {
		LogInternalError(w, errors.New("internal decoding error: "+err.Error()))
		return
	}

	// fmt.Println("New request", time.Now(), r.Method, uriDecoded)

	if uriDecoded == "/task/Циклическая ротация" && r.Method == http.MethodGet {
		getResp, err := http.Get("https://kuvaev-ituniversity.vps.elewise.com/tasks/" + url.PathEscape("Циклическая ротация"))
		// getResp, err := http.Get("http://localhost:3000/tasks")
		if err != nil {
			LogInternalError(w, errors.New("internal request error: "+err.Error()))
			return
		}

		//явно возвращаем ресурсы системе
		defer getResp.Body.Close()

		bytesssssssssssss, err := ioutil.ReadAll(getResp.Body)
		if err != nil {
			LogInternalError(w, errors.New("internal request byte parse error: "+err.Error()))
			return
		}

		//это парсер
		var args [][]json.RawMessage
		err = json.Unmarshal(bytesssssssssssss, &args)
		if err != nil {
			LogInternalError(w, errors.New("1 error: "+err.Error()))
		}
		var A = make([][]int, len(args))
		var K = make([]int, len(args))
		for i := 0; i < len(args); i++ {
			err = json.Unmarshal(args[i][0], &A[i])
			if err != nil {
				LogInternalError(w, errors.New("2 error: "+err.Error()))
			}
			err = json.Unmarshal(args[i][1], &K[i])
			if err != nil {
				LogInternalError(w, errors.New("internal request byte parse error: "+err.Error()))
			}
		}

		//это решение моим решателем
		var results = make([][]int, len(A))
		for i := range A {
			results[i] = internal.CyclicRotationSolve(A[i], K[i])
		}

		//отправка в сервис SOlutions
		values := BodyToSolutions{UserName: "test", Task: "Циклическая ротация", Results: Results{Payload: string(bytesssssssssssss), Results: results}}
		json_data, err := json.Marshal(values)
		if err != nil {
			LogInternalError(w, errors.New("3 error: "+err.Error()))
		}

		// postResp, err := http.Post("https://kuvaev-ituniversity.vps.elewise.com/tasks/solution", "application/json", bytes.NewBuffer(json_data))
		// if err != nil {
		// 	LogInternalError(w, errors.New("4 error: "+err.Error()))
		// }

		// bytesss, err := ioutil.ReadAll(postResp.Body)
		// if err != nil {
		// 	LogInternalError(w, errors.New("internal request byte parse error: "+err.Error()))
		// 	return
		// }
		fmt.Println("1")
		w.Write(json_data)
		fmt.Println(string(json_data))
		fmt.Println("2")

		return
	}

	if r.RequestURI == "/tasks" && r.Method == http.MethodGet {
		//обратываем запрос
		// if err != nil {
		// 	w.WriteHeader(http.StatusNotFound)
		// 	w.Write([]byte(err.Error()))
		// }
		// w.Write([]byte("tasks" + r.RequestURI))
		// var myint TaskPayload
		// myint = make(TaskPayload, 2)
		// myint[0] = []interface{}{[]int{1, 2, 3}, 999}
		// myint[1] = []interface{}{[]int{3, 2, 1}, 888}

		// marshalled, _ := json.Marshal(myint)
		// w.Write(marshalled)
		return
	}

	w.WriteHeader(http.StatusNotFound)
}
