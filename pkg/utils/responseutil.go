package utils

import (
	"encoding/json"
	"log"
	"net/http"
)

type ResponseDataStruct struct {
	Error              bool                   `json:"error"`
	Message            string                 `json:"message"`
	StatusCode         int                    `json:"statusCode"`
	ErrorCodeForClient string                 `json:"errorCodeForClient"`
	Data               map[string]interface{} `json:"data"`
}

func ResponseGenerator(w http.ResponseWriter, statusCode int, isError bool, errorCodeForClient string, responseDataMessage map[string]interface{}, message string) (responseJSON []byte) {
	responseData := ResponseDataStruct{
		Error:              isError,
		Message:            message,
		StatusCode:         statusCode,
		ErrorCodeForClient: errorCodeForClient,
		Data:               responseDataMessage,
	}

	responseJSON, err := json.Marshal(responseData)
	if err != nil {
		LogError("Failed to marshal output JSON", "Internal Server Error", 5, err, responseDataMessage)
		log.Fatal("Failed to marshal output JSON response util-", err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	w.Write(responseJSON)
	return
}
