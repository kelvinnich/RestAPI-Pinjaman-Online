package helper

import "strings"

type Response struct {
	Status  bool        `json:"status"`
	Message string      `json:"message"`
	Errors  interface{} `json:"errors"`
	Data    interface{} `json:"data"`
}


type EmptyObject struct {

}

func ResponseOK(status bool, message string, data interface{}) Response {
	reponses := Response{
		Status:  status,
		Message: message,
		Errors:  nil,
		Data:    data,
	}
	return reponses
}

func ErrorResponse(message string, err string, data interface{}) Response {
	splitErr := strings.Split(err, "\n")
	reponses := Response{
		Status:  false,
		Message: message,
		Errors:  splitErr,
		Data:    data,
	}
	return reponses
}