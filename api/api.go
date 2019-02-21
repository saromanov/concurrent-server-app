// Package api contains implementation of apis for endpoints
package api

import (
	"encoding/json"
	"net/http"
)

// QueryHandler implements query from shop
func QueryHandler(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	defer r.Body.Close()

	query := r.FormValue("query")
	results := getResults(query)

	response := result{
		Count: len(results),
		Data:  results,
	}

	result, err := json.Marshal(response)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Add("Content-Type", "application/json")
	w.Write(result)
}

// getResults returns results from shop
func getResults(query string) {

}
