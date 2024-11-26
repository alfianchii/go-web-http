package utils

import (
	"encoding/json"
	"fmt"
	"net/http"
	"web-http/dto"
)

func SetHeaderJson(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json")
}

func ResponseSetup(res http.ResponseWriter, req *http.Request) {
	SetHeaderJson(res)
	LogRequest(req)
}

func SendResponse(res http.ResponseWriter, msg string, status int, data interface{}) {
	response := dto.Response{
		Message: msg,
		Status: status,
		Data: data,
	}

	responseJson := ConvertToJson(response, res)
	
	ResponseWithJson(res, string(responseJson))
}

func ConvertToJson(data dto.Response, res http.ResponseWriter) []byte {
	dataJson, err := json.Marshal(data)
	if err != nil {
		http.Error(res, "Error converting response to JSON", http.StatusInternalServerError)
	}
	return dataJson
}

func ResponseWithName(res http.ResponseWriter, req *http.Request, msg string) {
	name := req.URL.Query().Get("name")
	if name == "" {
		name = "Guest"
	}

	message := msg + ", " + name + "!"
	SendResponse(res, message, http.StatusOK, nil)
}

func ResponseWithJson(res http.ResponseWriter, json string) {
	fmt.Fprint(res, json)
}