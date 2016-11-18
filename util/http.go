package util

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

const (
	ERROR   string = "ERROR"
	SUCCESS string = "OK"
)

// generic Json response struct
type Status struct {
	Status    string `json:"status"`
	Topic     string `json:"topic"`
	Message   string `json:"message"`
	Timestamp int64  `json:"timestamp"`
}

//
// Extract parameters from a request
//
func ExtractRequestParameter(r *http.Request, param string) string {
	return r.URL.Query().Get(param)
}

func ExtractRequestVariable(r *http.Request, param string) string {
	params := mux.Vars(r)
	return params[param]
}

//
// Helper functions for standard response scenarios
//

func OkResponse(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "text/plain; charset=utf-8") // normal header
	w.WriteHeader(http.StatusOK)
}

func SuccessResponse(w http.ResponseWriter, context, message string) {
	w.Header().Set("Content-Type", "application/json")

	status := Status{
		SUCCESS, context, message, Timestamp(),
	}

	response, err := json.Marshal(status)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	} else {
		w.Write(response)
	}
}

func JsonResponse(w http.ResponseWriter, model interface{}) {
	w.Header().Set("Content-Type", "application/json")

	response, err := json.Marshal(model)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	} else {
		w.WriteHeader(http.StatusOK)
		w.Write(response)
	}
}

func SimpleErrorResponse(w http.ResponseWriter, e error) {
	http.Error(w, e.Error(), http.StatusInternalServerError)
}

func GenericServerErrorResponse(w http.ResponseWriter, model interface{}) {
	w.Header().Set("Content-Type", "application/json")

	response, err := json.Marshal(model)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	} else {
		w.WriteHeader(http.StatusBadRequest)
		w.Write(response)
	}
}

func ErrorResponse(w http.ResponseWriter, context, message string) {
	w.Header().Set("Content-Type", "application/json")

	status := Status{
		ERROR, context, message, Timestamp(),
	}

	response, err := json.Marshal(status)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	} else {
		w.WriteHeader(http.StatusBadRequest)
		w.Write(response)
	}
}
