package utils

import (
	"log"
	"os"
)

type LogErrorData struct {
	TimeStamp          int64                  `json:"timeStamp"`
	ErrorMessage       string                 `json:"errorMessage"`
	ErrorCodeForClient string                 `json:"errorCodeForClient"`
	ErrorOrigin        string                 `json:"errorOrigin"`
	ErrorLevel         int                    `json:"errorLevel"`
	ErrorStack         error                  `json:"errorStack"`
	Data               map[string]interface{} `json:"data"`
}

func LogError(errorMessage string, errorCodeForClient string, errorLevel int, errorStack error, data map[string]interface{}) {
	errorOrigin := os.Getenv("SERVICE_NAME")
	LogErrorInfo := LogErrorData{
		TimeStamp:          GetCurrentTime().Unix(),
		ErrorMessage:       errorMessage,
		ErrorCodeForClient: errorCodeForClient,
		ErrorOrigin:        errorOrigin,
		ErrorLevel:         errorLevel,
		ErrorStack:         errorStack,
		Data:               data,
	}
	log.Println(LogErrorInfo)
}
