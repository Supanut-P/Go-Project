package api

import (
	"github.com/gorilla/mux"
)

func NewRouter() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/health", HealthCheck).Methods("GET")
	r.HandleFunc("/trigger/jenkins", TriggerJenkins).Methods("POST")
	r.HandleFunc("/logs/service/{name}", GetLogs).Methods("GET")
	r.HandleFunc("/docker/containers", ListContainers).Methods("GET")
	r.HandleFunc("/docker/container/start", StartContainer).Methods("POST")
	r.HandleFunc("/docker/container/stop", StopContainer).Methods("POST")
	return r
}
