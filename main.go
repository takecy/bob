package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/docopt/docopt-go"
	"github.com/takecy/bob/cli"
	"github.com/takecy/bob/config"

	d "github.com/tj/go-debug"
)

var debug = d.Debug("main")
var version = "0.0.1"

//command definition
const usage = `
  Bob is driver for Jenkins

  Usage:
    bob config
    bob env
    bob ping
    bob ls [--env env]
    bob ls <productname> [--env env]
    bob build <productname> <jobnumber> [--env env]
    bob build <productname> [--name <jobname>] [--env env]

  Options:
    -h --help           Print help.
    -v --version        Print version.
    --env env           Specify Environment. [default: local]
    --name jobname      Specify jobname, not jobnumber.
    --config configpath Specify custom config file path. [default: ./bob.yml]

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

	debug("[GOPATH] %s", gopath)
	debug("[args] %s", args)

	env := ""
	if _env, ok := args["--env"].(string); ok {
		env = _env
	}

	var configPath string
	if path, hasConfig := args["--config"].(string); hasConfig {
		configPath = path
	}

	configEnvVarPath := os.Getenv("BOB_CONFIG_PATH")
	if configPath == "" {
		if configEnvVarPath != "" {
			configPath = configEnvVarPath
		} else {
			configPath = "./bob.yml"
		}
	}

	debug("[BOB_CONFIG_PATH] %s", configEnvVarPath)
	debug("[env] %s", env)
	debug("[configPath] %s", configPath)

	bob, err := config.NewConfig(configPath)
	if err != nil {
		cli.Fatalf("read yaml error %s\n", err)
	}

	// commands switch
	switch {
	case args["config"].(bool):
		cli.Printf("Bob known Jenkins: \n%v\n", bob.ProductConfig)

	case args["env"].(bool):
		cli.ExecCommand("go", "env")

	case args["ping"].(bool):
		cli.Printf("PONG")

	case args["ls"].(bool):
		if prdName, hasPrdName := args["<productname>"].(string); hasPrdName {
			prdConf := bob.ProductConfig
			envConf, ok := prdConf[prdName]
			if ok {
				jenkinsConf := envConf[env]
				jobs, _ := cli.ListJobs(&jenkinsConf)

				for index, job := range jobs {
					fmt.Printf("%d:[%s]%s\n", index, job.Color, job.Name)
				}

				os.Exit(1)
			}
		} else {
			for prdName, envConf := range bob.ProductConfig {
				fmt.Printf("[%s]\n", prdName)

				for envName, jenkinsConf := range envConf {
					fmt.Printf("- %s\n", envName)
					fmt.Printf("  - %s\n", jenkinsConf.URL)
					fmt.Printf("  - %s\n", jenkinsConf.User)
					fmt.Printf("  - %s\n", jenkinsConf.Token)
				}
			}
			os.Exit(1)
		}

	case args["build"].(bool):
		if prdName, hasPrdName := args["<productname>"].(string); hasPrdName {
			prdConf := bob.ProductConfig
			envConf, ok := prdConf[prdName]

			if ok {
				jenkinsConf := envConf[env]
				jobs, _ := cli.ListJobs(&jenkinsConf)

				numStr, given := args["<jobnumber>"].(string)

				if !given {
					jobName, given := args["--name"].(string)
					if !given {
						cli.Fatalf("%s not exists.", "jobnumber")
						return
					}

					job, _ := cli.GetJob(&jenkinsConf, jobName)

					cli.Build(&jenkinsConf, job, nil)
					fmt.Println("Build Started: " + job.Name)

					os.Exit(1)
				} else {
					jobNum, _ := strconv.Atoi(numStr)

					job, _ := cli.SelectJob(jobs, jobNum)

					cli.Build(&jenkinsConf, job, nil)
					fmt.Println("Build Started: " + job.Name)

					os.Exit(1)
				}
			} else {
				cli.Fatalf("%s not exists.", "productname")
				return
			}
		} else {
			cli.Fatalf("%s required.", "productname")
			return
		}

	}

}
