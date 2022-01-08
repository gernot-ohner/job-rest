package main

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func runJobController() {

	r := mux.NewRouter()

	r.HandleFunc("/api/jobs", getJobs).Methods("GET")
	r.HandleFunc("/api/jobs/{id}", getJob).Methods("GET")
	r.HandleFunc("/api/jobs/{id}", createJob).Methods("POST")
	r.HandleFunc("/api/jobs/{id}", updateJob).Methods("PUT")
	r.HandleFunc("/api/jobs/{id}", deleteJob).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":8000", r))
}

func getJobs(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "application/json")
	json.NewEncoder(writer).Encode(jobs)
}

func getJob(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "application/json")
	id := mux.Vars(request)["id"]
	for _, job := range jobs {
		if job.Id == id {
			json.NewEncoder(writer).Encode(job)
			return
		}
	}
	json.NewEncoder(writer).Encode(&job{})
}

func createJob(writer http.ResponseWriter, request *http.Request) {

}

func updateJob(writer http.ResponseWriter, request *http.Request) {

}

func deleteJob(writer http.ResponseWriter, request *http.Request) {

}
