bob
===

Bob is the Driver of Jenkins.  
Jenkins is busy, and in everywhere.  (really a lot!)  
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
$bob ls --env dev
$bob ls <jobnumber>
$bob ls --name <jobname>
$bob build <jobnumber>
$bob build --name <jobname>
```

<br/>
### Quick Start
#### One Jenkins
If you need to drive for only one Jenkins.
```shell
$export BOB_JENKINS_URL=http://jenkins.example.com
$export BOB_JENKINS_API_TOKEN=absdefgh
$export BOB_PRODUCT_NAME=funky_pj
```

<br/>
#### More Jenkins
If you need to drive multiple Jenkins.  
Define your Jenkins config in `yaml` file.  

### Config yaml file
`yaml` file name default is `bob.yml`.  
and `yaml` file path is execute command directory.  
Should set environment variable, if you do not want to use default.
```shell
$export BOB_CONFIG_PATH=/usr/local/john.yml
```

#### yaml structure
```yaml
//TODO
blogs:
  dev:
    dev-jenkins.blogs.com
  stg:
    stg-jenkins.blogs.com
  prd:
    jenkins.blogs.com
twitclone:
  sbx:
    sbx-jenkins.twc.com
  prd:
    jenkins.twc.com
```

<br/>
Development
----
1. Fork to your repository.
1. ``` $git clone <your repository url>```
1. ```$make build```
1. ```$./bob ping```
1. modify files and push.
1. Pull Request !

<br/>
License
----
MIT
