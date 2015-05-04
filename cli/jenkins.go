package cli

import (
	"net/url"

	"github.com/takecy/bob/entity"
	gojenkins "github.com/yosida95/golang-jenkins"
)

// ListJobs jenkins jobs
func ListJobs(conf *entity.JenkinsConfig) (jobs []gojenkins.Job, err error) {
	auth := &gojenkins.Auth{
		Username: conf.User,
		ApiToken: conf.Token,
	}
	jenkins := gojenkins.NewJenkins(auth, conf.URL)
	jobs, err = jenkins.GetJobs()

	if err != nil {
		Fatalf("error exec command: %s", err)
		return
	}

	return
}

// GetJob specify jenkins job
func GetJob(conf *entity.JenkinsConfig, jobName string) (job gojenkins.Job, err error) {
	auth := &gojenkins.Auth{
		Username: conf.User,
		ApiToken: conf.Token,
	}
	jenkins := gojenkins.NewJenkins(auth, conf.URL)
	job, err = jenkins.GetJob(jobName)

	if err != nil {
		Fatalf("error exec command: %s", err)
		return
	}

	return
}

// Build specify jenkins job
func Build(conf *entity.JenkinsConfig, job gojenkins.Job, params url.Values) (err error) {
	auth := &gojenkins.Auth{
		Username: conf.User,
		ApiToken: conf.Token,
	}
	jenkins := gojenkins.NewJenkins(auth, conf.URL)
	err = jenkins.Build(job, params)

	if err != nil {
		Fatalf("error exec command: %s", err)
		return
	}

	return
}

// SelectJob from jobs slice
func SelectJob(jobs []gojenkins.Job, number int) (job gojenkins.Job, err error) {
	for index, job := range jobs {
		if index == number {
			return job, nil
		}
	}

	Fatalf("error exec command: %s", err)
	return
}
