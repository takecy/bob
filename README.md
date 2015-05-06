bob
===

Who is Bob?
----

Bob is the Driver of Jenkins.  
Jenkins is so busy, and in everywhere.  (really a lot!)  
<br/>
![jenkins](./img/s_jenkins.png)
![jenkins](./img/s_jenkins.png)
![jenkins](./img/s_jenkins.png)
![jenkins](./img/s_jenkins.png)
![jenkins](./img/s_jenkins.png)
![jenkins](./img/s_jenkins.png)
.....  

<br/>
Some product has Trouble?  
My phone is receiving alerts now...  
Responsible Jenkins Where is it!?  

Let's listen to Bob !

<br/>

Features
----
* List jobs and Build job on command line.
* Managed Jenkins hosts at product name.
* Change easy environment.

<br/>

Installation
----
#### via go-get
bob is written by golang.  
```shell
$go version
go version go1.4.2 darwin/amd64
```  
installed in your $GOPATH/src
```shell
$go get github.com/takecy/bob
```

#### via binary
not supported yet.
```
//TODO
```

<br/>

Usage
---
print help
```shell
$bob -h
```

exmaple
```shell
$bob ping
$bob ls | grep dev
$bob ls <productname>
$bob ls <jobnumber>
$bob ls --name <jobname>
$bob build <jobnumber>
$bob build --name <jobname>
```

<br/>
### Configuration
If you need to drive multiple Jenkins.  
Define your Jenkins config in `yaml` file.  

#### Config yaml file
`yaml` file name default is `bob.yml`.  
and `yaml` file path is execute command directory.  
Should set environment variable, if you do not want to use default.
```shell
$export BOB_CONFIG_PATH=/usr/local/john.yml
```
or `--config` option.

##### yaml structure
```yaml
blogs:                         // product name (required)
  dev:                         // environment  (required)
    url: dev-jenkins.blogs.com // jenkins URL  (required)
    user: john                 // jenkins user name (optional)
    token: 12345               // jenkins API token (optional)
  stg:
    url: stg-jenkins.blogs.com
    token: abc
  prd:
    url: jenkins.blogs.com
    user: noris
twitclone:
  sbx:
    url: sbx-jenkins.twc.com
    user: michael
  prd:
    url: jenkins.twc.com
    user: lebron
    token: james_token
```

<br/>
Development
----
1. Fork to your repository.
1. modify files and push.
1. Pull Request !

```shell
$godep restore
$make build
$DEBUG=* ./bob ping
```

<br/>
License
----
MIT
