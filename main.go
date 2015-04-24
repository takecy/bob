package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/docopt/docopt-go"
	"github.com/takecy/bob/cli"
	"github.com/takecy/bob/config"
)

var version = "0.0.1"

//command definition
const usage = `
  command-line tool for jenkins

  Usage:
    bob ping [--debug]
    bob ls [--env env]
    bob ls <jobnumber> [--env env]
    bob ls [--name <jobname>] [--env env]
    bob build <jobnumber> [--env env]
    bob build [--name <jobname>] [--env env]

  Options:
    --debug           Print debug log.
    -h --help         Print help.
    -v --version      Print version.
    -i identityfile   ssh identityfile path. [default: ~/.ssh/liverpool.pem]
    --env env         Specify Environment (local|dev|stg). [default: local]
    --name jobname    Specify jobname, not jobnumber.

  Examples:
    $bob ls
    $bob ls --env dev
    $bob ls 30
    $bob build 30
		`

//main
func main() {
	gopath := os.Getenv("GOPATH")
	if gopath == "" {
		cli.Fatalf("can't get $GOPATH")
	}

	configFilePath := os.Getenv("BOB_CONFIG_PATH")
	if configFilePath == "" {
		configFilePath = "bob.yml"
	}

	bob, _ := config.NewConfig(configFilePath)

	args, err := docopt.Parse(usage, nil, true, version, false)
	if err != nil {
		cli.Fatalf("error parsing args: %s", err)
	}

	if args["--debug"].(bool) {
		fmt.Printf("$GOAPTH -> %s\n", gopath)
		fmt.Println("args", args)
	}

	switch {
	case args["ping"].(bool):
		fmt.Println("PONG")

	case args["ls"].(bool):
		if numberStr, hasName := args["<jobnumber>"].(string); hasName {
			number, err := strconv.Atoi(numberStr)
			if err != nil {
				cli.Fatalf("bad number %s", numberStr)
			}

			jobs, _ := cli.ListJobs(bob)
			job, _ := cli.SelectJob(bob, jobs, number)
			fmt.Printf("[%s]%s\n", job.Color, job.Name)
		} else {
			jobs, _ := cli.ListJobs(bob)
			for i, job := range jobs {
				fmt.Printf("[%d][%s]%s\n", i, job.Color, job.Name)
			}
		}

	case args["build"].(bool):
		if numberStr, hasName := args["<jobnumber>"].(string); hasName {
			number, err := strconv.Atoi(numberStr)
			if err != nil {
				cli.Fatalf("bad number %s", numberStr)
				return
			}

			jobs, _ := cli.ListJobs(bob)
			job, _ := cli.SelectJob(bob, jobs, number)
			cli.Build(bob, job, nil)
			fmt.Println("Build Started: " + job.Name)

		} else if jobName, hasName := args["--name"].(string); hasName {
			job, err := cli.GetJob(bob, jobName)
			if err != nil {
				cli.Fatalf("bad number %s", numberStr)
				return
			}

			cli.Build(bob, job, nil)
			fmt.Println("Build Started: " + job.Name)
		} else {
			cli.Fatalf("jobnumber or jobname is required. %s", "build command")
			return
		}

	}

}
