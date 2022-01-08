package main

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"log"
	"math/rand"
	"net/http"
	"strconv"
)

func runJobController() {

	r := mux.NewRouter()

	r.HandleFunc("/api/jobs", getJobs).Methods("GET")
	r.HandleFunc("/api/jobs/{id}", getJob).Methods("GET")
	r.HandleFunc("/api/jobs", createJob).Methods("POST")
	r.HandleFunc("/api/jobs/{id}", updateJob).Methods("PUT")
	r.HandleFunc("/api/jobs/{id}", deleteJob).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":8000", r))
}

func getJobs(writer http.ResponseWriter, _ *http.Request) {
	writer.Header().Set("Content-Type", "application/json")
	log.Fatal(json.NewEncoder(writer).Encode(jobs))
}

func getJob(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "application/json")
	id := mux.Vars(request)["id"]
	for _, job := range jobs {
		if job.Id == id {
			log.Fatal(json.NewEncoder(writer).Encode(job))
			return
		}
	}
	log.Fatal(json.NewEncoder(writer).Encode(&Job{}))
}

func createJob(writer http.ResponseWriter, request *http.Request) {
	var job Job
	_ = json.NewDecoder(request.Body).Decode(&job)
	job.Id = strconv.Itoa(rand.Intn(100000))
	jobs = append(jobs, job)
	log.Fatal(json.NewEncoder(writer).Encode(job))
}

func updateJob(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "application/json")
	id := mux.Vars(request)["id"]

	var newjob Job
	_ = json.NewDecoder(request.Body).Decode(&newjob)
	for index, job := range jobs {
		if job.Id == id {
			jobs[index] = newjob
		}
	}

	log.Fatal(json.NewEncoder(writer).Encode(jobs))
}

func deleteJob(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "application/json")
	id := mux.Vars(request)["id"]
	for index, job := range jobs {
		if job.Id == id {
			jobs = append(jobs[:index], jobs[index+1:]...)
			break
		}
	}

	log.Fatal(json.NewEncoder(writer).Encode(jobs))
}
