package response

import (
	"encoding/json"
	"gocrawler-web-sample/shared/times"
	"log"
	"net/http"
)

// HTTPResponse func
func HTTPResponse(w http.ResponseWriter, data interface{}, statusCode int) {
	responseTemplate := ResponseStruct{
		Data:   data,
		TimeIn: times.Now(times.TimeGmt, times.TimeFormat),
	}
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.WriteHeader(statusCode)
	resp, err := json.Marshal(responseTemplate)
	if err != nil {
		log.Fatal(err)
	}
	_, _ = w.Write(resp)
}

// SendSuccessResponse function with HTTPResponse
func SendSuccessResponse(w http.ResponseWriter, data interface{}) {
	HTTPResponse(w, data, http.StatusOK)
}

// SendErrorResponse function with HTTPResponse
func SendErrorResponse(w http.ResponseWriter, data interface{}, statusCode int) {
	HTTPResponse(w, data, statusCode)
}
