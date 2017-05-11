package service

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/manyminds/api2go/jsonapi"

	"github.com/romatr/interestcalculator/formula"
	"github.com/romatr/interestcalculator/model"
)

func handleCalculate(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	query := r.URL.Query()

	name := vars["name"]

	amount, err := strconv.ParseFloat(query.Get("amount"), 64)
	if err != nil {
		writeErrorResponse(w, http.StatusBadRequest, "PARAM_MISSING", "Required parameter 'amount' is missing")
		return
	}

	interestRate, err := strconv.ParseFloat(query.Get("interest_rate"), 64)
	if err != nil {
		writeErrorResponse(w, http.StatusBadRequest, "PARAM_MISSING", "Required parameter 'interest_rate' is missing")
		return
	}

	termInMonths, err := strconv.ParseInt(query.Get("term_in_months"), 10, 64)
	if err != nil {
		writeErrorResponse(w, http.StatusBadRequest, "PARAM_MISSING", "Required parameter 'term_in_months' is missing")
		return
	}

	interest := formula.CalculateInterest(name, amount, interestRate, termInMonths)
	writeDataResponse(w, http.StatusOK, model.InterestResponse{ID: name, Interest: interest})
}

func writeDataResponse(w http.ResponseWriter, status int, dataModel model.InterestResponse) {
	jsonResponse, _ := jsonapi.Marshal(dataModel)

	writeJsonResponse(w, status, jsonResponse)
}

func writeErrorResponse(w http.ResponseWriter, status int, code string, title string) {
	error := model.Error{Code: code, Title: title}
	errorResponse, _ := json.Marshal(model.ErrorResponse{Errors: []model.Error{error}})

	writeJsonResponse(w, status, errorResponse)
}

func writeJsonResponse(w http.ResponseWriter, status int, data []byte) {
	w.Header().Set("Content-Type", "application/vnd.api+json")
	w.Header().Set("Content-Length", strconv.Itoa(len(data)))
	w.WriteHeader(status)
	w.Write(data)
}
