package services

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
)

type CrumbResponse struct {
	Crumb             string `json:"crumb"`
	CrumbRequestField string `json:"crumbRequestField"`
}

func TriggerJenkinsJob(job string) error {
	jenkinsURL := os.Getenv("JENKINS_URL")
	user := os.Getenv("JENKINS_USER")
	token := os.Getenv("JENKINS_TOKEN")

	// ดึง crumb token
	crumbURL := fmt.Sprintf("%s/crumbIssuer/api/json", jenkinsURL)
	req, _ := http.NewRequest("GET", crumbURL, nil)
	req.SetBasicAuth(user, token)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return fmt.Errorf("error getting crumb: %w", err)
	}
	defer resp.Body.Close()

	body, _ := io.ReadAll(resp.Body)
	var crumbRes CrumbResponse
	json.Unmarshal(body, &crumbRes)

	// ใช้ crumb ใน request จริง
	buildURL := fmt.Sprintf("%s/job/%s/build", jenkinsURL, job)
	reqBuild, _ := http.NewRequest("POST", buildURL, nil)
	reqBuild.SetBasicAuth(user, token)
	reqBuild.Header.Set(crumbRes.CrumbRequestField, crumbRes.Crumb)

	respBuild, err := client.Do(reqBuild)
	if err != nil || respBuild.StatusCode >= 400 {
		return fmt.Errorf("build trigger failed: %v", respBuild.Status)
	}

	return nil
}
