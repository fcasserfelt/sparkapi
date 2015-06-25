package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/thoas/stats"
	"net/http"
)

type GraphResponse struct {
	Node string `json:"node"`
	Link string `json:"link"`
}

func ApiRouter(statsMiddleware *stats.Stats) *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/api", ApiHandler)
	router.HandleFunc("/api/graph", GraphListHandler)
	router.HandleFunc("/api/graph/{id}", GraphHandler)
	router.HandleFunc("/api/stats", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		stats := statsMiddleware.Data()
		b, _ := json.Marshal(stats)
		w.Write(b)
	})
	return router
}

func ApiHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "I'm Api")
}

func GraphListHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "All my graphs")
}

func GraphHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	key := vars["id"]

	response := &GraphResponse{Node: "Ny node: " + key, Link: "My link"}

	submitJsonResponse(w, response)
}

func submitJsonResponse(w http.ResponseWriter, resp *GraphResponse) {
	w.Header().Set("Content-Type", "application/json")

	jsonResponse, jsonErr := ToJSON(resp)
	if jsonErr != nil {
		http.Error(w, "{\"error\": \"Unknown Error\"}", 500)
		return
	}
	w.Write([]byte(jsonResponse))
}
