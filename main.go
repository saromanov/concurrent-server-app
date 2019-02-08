package main

import (
	"fmt"
	"net/http"
	"os"

)

func main() {
	http.HandleFunc("/api/analysis", api.Analysis)
	http.HandleFunc("/api/query", api.Query)

	api.Start()

	address := fmt.Sprintf(":%s", os.Getenv("ADDRESS"))
	http.ListenAndServe(port, nil)
}