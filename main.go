package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/docopt/docopt-go"
	"github.com/takecy/bob/cli"
)

var version = "0.0.1"

//command definition
const usage = `
  command-line tool for jenkins

  Usage:
    bob ping [--debug]
    bob jenkins ls
    bob jenkins ls <jobnumber>
    bob jenkins build <jobnumber>

  Options:
    --debug           Print debug log.
    -h --help         Print help.
    -v --version      Print version.
    -i identityfile   ssh identityfile path. [default: ~/.ssh/liverpool.pem]
    -e env            Specify Environment (local|dev|stg). [default: local]

  Examples:
    jenkins
    $bob jenkins ls
    $bob jenkins ls 30
    $bob jenkins build 30
		`

//main
func main() {
	gopath := os.Getenv("GOPATH")
	if gopath == "" {
		cli.Fatalf("can't get $GOPATH")
	}

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

	case args["jenkins"].(bool) && args["list"].(bool):
		if numberStr, hasName := args["<jobnumber>"].(string); hasName {
			number, err := strconv.Atoi(numberStr)
			if err != nil {
				cli.Fatalf("bad number %s", numberStr)
			}

			jobs, _ := cli.ListJobs()
			job, _ := cli.SelectJob(jobs, number)
			fmt.Printf("[%s]%s\n", job.Color, job.Name)
		} else {
			jobs, _ := cli.ListJobs()
			for i, job := range jobs {
				fmt.Printf("[%d][%s]%s\n", i, job.Color, job.Name)
			}
		}

	case args["jenkins"].(bool) && args["build"].(bool):
		fmt.Println("jenkins build")
		if numberStr, hasName := args["<jobnumber>"].(string); hasName {
			number, err := strconv.Atoi(numberStr)
			if err != nil {
				cli.Fatalf("bad number %s", numberStr)
			}

			jobs, _ := cli.ListJobs()
			job, _ := cli.SelectJob(jobs, number)
			cli.Build(job, nil)
			fmt.Println("Build Started: " + job.Name)
		}

	}

}
