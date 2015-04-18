bob
===

Bob is the Driver of Jenkins.  
Jenkins is busy, and in everywhere.  (really a lot!)  
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
via ```go-get```
```
$go get github.com/takecy/bob
```
via binary
```
//TODO
```

<br/>

Usage
---
```
$bob -h
$bob -v
$bob ping
$bob config
$bob ls <productname>
$bob ls --env dev
$bob ls <jobnumber>
$bob ls --name <jobname>
$bob build <jobnumber>
$bob build --name <jobname>
```

### Config yaml file
yml file name must be ```bob.yml```  

#### example
```
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
License
----
MIT
