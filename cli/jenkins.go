package cli

import (
	"net/url"

	gojenkins "github.com/yosida95/golang-jenkins"
)

var auth = &gojenkins.Auth{
	Username: "",
	ApiToken: "",
}

// Bob inclueds Jenkins info
type Bob struct {
	Jenkins    gojenkins.Auth
	JenkinsURL string
}

const jenkinsURL = "http://jenkins.awa.io"

// ListJobs jenkins jobs
func ListJobs(bob Bob) (jobs []gojenkins.Job, err error) {
	jenkins := gojenkins.NewJenkins(auth, jenkinsURL)
	jobs, err = jenkins.GetJobs()

	if err != nil {
		Fatalf("error exec command: %s", err)
		return
	}

	return
}

// GetJob specify jenkins job
func GetJob(bob Bob, jobName string) (job gojenkins.Job, err error) {
	jenkins := gojenkins.NewJenkins(auth, jenkinsURL)
	job, err = jenkins.GetJob(jobName)

	if err != nil {
		Fatalf("error exec command: %s", err)
		return
	}

	return
}

// Build specify jenkins job
func Build(bob Bob, job gojenkins.Job, params url.Values) (err error) {
	jenkins := gojenkins.NewJenkins(auth, jenkinsURL)

	err = jenkins.Build(job, params)

	if err != nil {
		Fatalf("error exec command: %s", err)
		return
	}

	return
}

// SelectJob from jobs slice
func SelectJob(bob Bob, jobs []gojenkins.Job, number int) (job gojenkins.Job, err error) {
	for index, job := range jobs {
		if index == number {
			return job, nil
		}
	}

	Fatalf("error exec command: %s", err)
	return
}
