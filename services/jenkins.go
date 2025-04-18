package services

import (
	"fmt"
	"net/http"
	"os"
)

func TriggerJenkinsJob(job string) error {
	jenkinsURL := os.Getenv("JENKINS_URL")
	url := fmt.Sprintf("%s/job/%s/build", jenkinsURL, job)
	resp, err := http.Post(url, "application/json", nil)
	if err != nil || resp.StatusCode != http.StatusCreated {
		return err
	}
	return nil
}
