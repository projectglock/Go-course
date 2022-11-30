package tasksHandler

import (
	"encoding/json"
	"errors"
	"final_task/service"
	"final_task/tasksSolvers"
	"net/http"
)

// BodyRequest - тело для POST запроса
type BodyRequest struct {
	UserName string      `json:"user_name"` //имя юзера указанное в тг
	Task     string      `json:"task"`      //имя задачи
	Results  BodyResults `json:"results"`
}

// BodyResults - часть POST запроса
type BodyResults struct {
	Payload [][][]int `json:"payload"` //данные полученные для решения задачи
	Results []int     `json:"results"` //результаты полученные при решении задачи
}

func TasksHandler(w http.ResponseWriter, urlEndpoint string) {
	//Задачку "Циклическая ротация" сразу уводим в отдельную функцию
	if urlEndpoint == "Циклическая ротация" {
		CyclicRotationHandler(w)
		return
	}

	//получаем данные для задачи
	respBytes, err := service.GetRequest(urlEndpoint)
	if err != nil {
		service.LogInternalError(w, err)
		return
	}

	//парсим входные данные
	var parsedData [][][]int
	err = json.Unmarshal(respBytes, &parsedData)
	if err != nil {
		service.LogInternalError(w, errors.New("internal request byte parse error: "+err.Error()))
		return
	}

	//запуск решателей задач
	//из GET запроса приходит массив лишний раз обёрнутый в массив, поэтому приходится обращаться к 0 индексу parsedData[i][0]
	var results = make([]int, len(parsedData))
	for i := range parsedData {
		switch urlEndpoint {
		case "Чудные вхождения в массив":
			results[i] = tasksSolvers.WondOccurrencesArraySolve(parsedData[i][0])
		case "Проверка последовательности":
			results[i] = tasksSolvers.CheckSequenceSolve(parsedData[i][0])
		case "Поиск отсутствующего элемента":
			results[i] = tasksSolvers.FindMissElemSolve(parsedData[i][0])
		}
	}

	//подготовка данных для POST запроса
	values := BodyRequest{UserName: "Коршунов А (pojectglock)", Task: urlEndpoint, Results: BodyResults{Payload: parsedData, Results: results}}
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
