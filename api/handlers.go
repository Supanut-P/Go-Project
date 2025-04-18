package api

import (
	"devops-toolbelt/services"
	"net/http"

	"github.com/gorilla/mux"
)

func HealthCheck(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("OK"))
}

func TriggerJenkins(w http.ResponseWriter, r *http.Request) {
	job := r.URL.Query().Get("job")
	if job == "" {
		http.Error(w, "Missing job name", http.StatusBadRequest)
		return
	}

	err := services.TriggerJenkinsJob(job)
	if err != nil {
		http.Error(w, "Failed to trigger job", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Job triggered successfully"))
}

func GetLogs(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	name := vars["name"]
	logs := services.GetServiceLogs(name)
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(logs))
}

func ListContainers(w http.ResponseWriter, r *http.Request) {
	containers := services.ListRunningContainers()
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(containers))
}

func StartContainer(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name")
	msg := services.StartContainer(name)
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(msg))
}

func StopContainer(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name")
	msg := services.StopContainer(name)
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(msg))
}
