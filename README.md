Distributed system for Crawl using by Golang
=========================
[![Build Status](https://travis-ci.org/zjucx/golang-webserver.svg?branch=master
)](http://120.27.39.169:8080/home)
[![Yii2](https://img.shields.io/badge/PoweredBy-ZjuCx-brightgreen.svg?style=flat)](http://120.27.39.169:8080/home)

Requirements
-----

* Docker(1.1x)

* Golang(1.6)

What you need is only the Golang environment and docker machine!


Using
------------
```
<!--  Prepare redis servre and containers for worker  --!>
git clone https://github.com/zjucx/redismq.git
cd redismq
go get
// for master
go run main.go master redisserverip:port masterserverip:port
// for workers
go run main.go worker masterserverip:port workerserverip:port
```

To Do List
----------

- Add scrawl for this system

Update
-----------------
```
git pull
```

Screenshots
-----------

### demo show
#### design
![](https://github.com/zjucx/redismq/blob/master/docs/distributeredis.bmp)
#### master
![](https://github.com/zjucx/redismq/blob/master/docs/master.png)
#### workers
![](https://github.com/zjucx/redismq/blob/master/docs/worker.png)
![](https://github.com/zjucx/redismq/blob/master/docs/contains.png)

Discussing
----------
- [submit issue](https://github.com/zjucx/redismq/issues/new)
- email: zju.chx@gmail.com
