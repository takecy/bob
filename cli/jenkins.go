package cli

import (
	"net/url"

	"github.com/takecy/bob/entity"
	gojenkins "github.com/yosida95/golang-jenkins"

	d "github.com/tj/go-debug"
)

var debug = d.Debug("cli")

// ListJobs jenkins jobs
func ListJobs(conf *entity.JenkinsConfig) (jobs []gojenkins.Job, err error) {
	debug("[ListJobs]conf", *conf)

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
	debug("[GetJob]conf", *conf)
	debug("[GetJob]jobName", jobName)

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
	debug("[Build]conf", *conf)
	debug("[Build]job", job)
	debug("[Build]params", params)

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
	debug("[SelectJob]conf", jobs)
	debug("[SelectJob]number", number)

	for index, job := range jobs {
		if index == number {
			return job, nil
		}
	}

	Fatalf("cant selected job from number : %s", err)
	return
}
