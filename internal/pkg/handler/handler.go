package handler

import (
	"encoding/json"
	"net/http"

	"github.com/meziaris/gofinance/internal/pkg/validator"
)

type ResponseBody struct {
	Status  string      `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

// Handle response error
func ResponseError(w http.ResponseWriter, statusCode int, message string) {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	encoder := json.NewEncoder(w)
	resp := ResponseBody{
		Status:  "error",
		Message: message,
	}
	_ = encoder.Encode(resp)
}

// Handle response success
func ResponseSuccess(w http.ResponseWriter, statusCode int, message string, data interface{}) {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	encoder := json.NewEncoder(w)
	resp := ResponseBody{
		Status:  "success",
		Message: message,
		Data:    data,
	}
	_ = encoder.Encode(resp)
}

// Parse request data & validate struct
func BindAndCheck(w http.ResponseWriter, r *http.Request, data interface{}) bool {
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(data); err != nil {
		ResponseError(w, http.StatusUnprocessableEntity, err.Error())
		return true
	}

	valid := validator.Check(data)
	if !valid {
		ResponseError(w, http.StatusUnprocessableEntity, "request format is not valid")
		return true
	}

	return false
}
