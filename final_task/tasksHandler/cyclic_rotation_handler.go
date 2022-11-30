// Для задачи Циклическая ротация отдельный handler потому что у неё входные данные отличаются от других задач
// следовательно их надо парсить и выводить по другому
// не очень хорошо, что часть функционала тут повторяется, не придумал как лучше
package tasksHandler

import (
	"encoding/json"
	"errors"
	"final_task/service"
	"final_task/tasksSolvers"
	"net/http"
)

// BodyForSolutions - тело для POST запроса
type BodyRequestCyclicRotation struct {
	UserName string                `json:"user_name"` //имя юзера указанное в тг
	Task     string                `json:"task"`      //имя задачи
	Results  ResultsCyclicRotation `json:"results"`
}

// Results - часть POST запроса
type ResultsCyclicRotation struct {
	Payload [][]json.RawMessage `json:"payload"` //данные полученные для решения задачи
	Results [][]int             `json:"results"` //результаты полученные при решении задачи
}

func CyclicRotationHandler(w http.ResponseWriter) {
	//получаем данные для задачи
	respBytes, err := service.GetRequest("Циклическая ротация")
	if err != nil {
		service.LogInternalError(w, err)
		return
	}

	//парсим входные данные в две переменные A и K
	var parsedData [][]json.RawMessage
	err = json.Unmarshal(respBytes, &parsedData)
	if err != nil {
		service.LogInternalError(w, errors.New("internal request byte parse error: "+err.Error()))
		return
	}
	//A - это слайс с исходными массивами из ответа
	//K - это массив чисел-номеров итераций для ротации
	var A = make([][]int, len(parsedData))
	var K = make([]int, len(parsedData))
	for i := 0; i < len(parsedData); i++ {
		err = json.Unmarshal(parsedData[i][0], &A[i])
		if err != nil {
			service.LogInternalError(w, errors.New("internal request byte parse error: "+err.Error()))
		}
		err = json.Unmarshal(parsedData[i][1], &K[i])
		if err != nil {
			service.LogInternalError(w, errors.New("internal request byte parse error: "+err.Error()))
		}
	}

	//запуск решателя задачи
	var results = make([][]int, len(A))
	for i := range A {
		results[i] = tasksSolvers.CyclicRotationSolve(A[i], K[i])
	}

	//подготовка данных для POST запроса
	values := BodyRequestCyclicRotation{UserName: "Коршунов А (pojectglock)", Task: "Циклическая ротация", Results: ResultsCyclicRotation{Payload: parsedData, Results: results}}
	json_data, err := json.Marshal(values)
	if err != nil {
		service.LogInternalError(w, errors.New("internal Marshal error: "+err.Error()))
		return
	}

	//отправка POST запроса в сервис Solutions
	resp, err := service.PostRequest(&json_data)
	if err != nil {
		service.LogInternalError(w, errors.New("internal request byte parse error: "+err.Error()))
		return
	}

	w.Write(resp)
	return
}
