package service

import (
	"bytes"
	"errors"
	"io/ioutil"
	"net/http"
	"net/url"
)

func GetRequest(endpiont string) ([]byte, error) {
	resp, err := http.Get("https://kuvaev-ituniversity.vps.elewise.com/tasks/" + url.PathEscape(endpiont))
	if err != nil {
		return nil, errors.New("internal request error: " + err.Error())
	}

	//явно возвращаем ресурсы системе
	defer resp.Body.Close()

	respBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, errors.New("iinternal request byte parse error: " + err.Error())
	}
	return respBytes, nil
}

func PostRequest(json_data *[]byte) ([]byte, error) {
	resp, err := http.Post("https://kuvaev-ituniversity.vps.elewise.com/tasks/solution", "application/json", bytes.NewBuffer(*json_data))
	if err != nil {
		return nil, errors.New("internal request error: " + err.Error())
	}

	//явно возвращаем ресурсы системе
	defer resp.Body.Close()

	respBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, errors.New("iinternal request byte parse error: " + err.Error())
	}
	return respBytes, nil
}
