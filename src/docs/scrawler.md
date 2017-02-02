## Distributed system for Crawl using by Golang

[![Build Status](https://travis-ci.org/zjucx/golang-webserver.svg?branch=master
)](http://120.27.39.169:8080/home)
[![Yii2](https://img.shields.io/badge/PoweredBy-ZjuCx-brightgreen.svg?style=flat)](http://120.27.39.169:8080/home)

### Introduction
使用golang开发的[分布式爬虫系统](https://github.com/zjucx/DistributedCrawler.git)，主要分为3个模块:[分布式框架](src/docs/framework.md)、[数据管理](src/docs/model.md)和[爬虫部分](src/docs/scrawler.md)。目录结构如下:
```
├── main.go
└── scrawler           ------定义了数据库模型，用于与数据库交互
    ├── sinaLogin.go   ------模拟登陆模块，工程中实现了新浪微博的模拟登陆
    ├── scrawler.go    ------爬虫模块的入口，将接口暴漏于分布式模块
    ├── scheduler.go   ------爬虫的调度器，由于对master分发的url任务的预处理
    ├── downloader.go  ------爬虫的下载器，管理多个下载任务的同步等操作
    ├── spiders.go     ------爬虫的数据提取，用于提取resp的url和想要爬取的数据
    ├── pipeline.go    ------url和目的数据的持久化操作
    ├── request.go     ------封装的request请求
    └── utils.go       ------爬虫的辅助类
```
#### design
![](https://github.com/zjucx/redismq/blob/master/docs/distributeredis.bmp)
### [爬虫部分](src/docs/scrawler.md)
```
用golang重写了新浪微博的模拟登陆,期间尼玛的各种坑。爬虫框架部分采用了python的Scrapy组件设计(实际上是用了Scrapy框架个组件的名字)。
```
#### 模拟登陆
```
参考[模拟登陆](https://zhuanlan.zhihu.com/p/23064000)链接
重要的事情说三遍:分析登陆原理、原理、原理
爬虫的重要原则:能爬手机爬手机,为vuejs布个道(新浪微博手机端用vue重写了)
```
#### 爬虫框架
参考官网框架图:
![](https://doc.scrapy.org/en/1.3/_images/scrapy_architecture_02.png)
### Discussing
- [submit issue](https://github.com/zjucx/DistributedCrawler/issues/new)
- email: 862575451@qq.com
