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
  Bob is driver for Jenkins

  Usage:
    bob config
    bob env
    bob ping [--debug]
    bob ls [--env env]
    bob ls <productname>
    bob build <jobnumber> [--env env]
    bob build [--name <jobname>] [--env env]

  Options:
    --debug             Print debug log.
    -h --help           Print help.
    -v --version        Print version.
    --env env           Specify Environment. [default: local]
    --name jobname      Specify jobname, not jobnumber.
    --config configpath Specify custom config file path.[default: ./bob.yml]

  Examples:
    $bob ls --env dev
    $bob ls hoge_product
    $bob build 30
		`

//main
func main() {

	gopath := os.Getenv("GOPATH")
	if gopath == "" {
		cli.Fatalf("can't get $GOPATH")
	}

	// parse command line arguments
	args, err := docopt.Parse(usage, nil, true, version, false)
	if err != nil {
		cli.Fatalf("error parsing args: %s", err)
	}

	if args["--debug"].(bool) {
		fmt.Printf("$GOAPTH -> %s\n", gopath)
		fmt.Println("args", args)
	}

	var configPath string
	if path, hasConfig := args["--config"].(string); hasConfig {
		configPath = path
	}

	if configPath == "" {
		configEnvVarPath := os.Getenv("BOB_CONFIG_PATH")
		if configEnvVarPath != "" {
			configPath = configEnvVarPath
		} else {
			configPath = "bob.yml"
		}
	}

	bob, err := config.NewConfig(configPath)
	if err != nil {
		cli.Fatalf("read yaml error %s\n", err)
	}

	// commands switch
	switch {
	case args["config"].(bool):
		fmt.Printf("Bob known Jenkins: \n%v\n", *bob.ProductConfig)

	case args["env"].(bool):
		cli.ExecCommand("go", "env")

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
