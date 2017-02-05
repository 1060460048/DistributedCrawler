## Distributed system for Crawl using by Golang -- Framework
### Introduction
使用golang开发的[分布式爬虫系统](https://github.com/zjucx/DistributedCrawler.git)，主要分为3个模块:[分布式框架](src/docs/framework.md)、[数据管理](src/docs/model.md)和[爬虫部分](src/docs/scrawler.md)。分布式框架部分目录结构如下:
```
└── distribute    
    ├── common.go      ------分布式系统的辅助类的定义等
    ├── master.go      ------分布式系统的master节点，任务的分发调度
    └── worker.go      ------分布式系统的worker节点，接受master的任务
```
### [分布式框架](src/docs/framework.md)
#### master
```
主要分为两部分功能:(1)使用redis做url缓存;(2)分发url到合适的worker。
```
##### url缓存实现
```
关键技术点:(1)redis缓存在合适的时机从mongodb获取url;(2)如何处理生产者和消费者的关系。
(一)合适的时机:
1)worker节点操作完成--此时说明已经有新的url存入mongodb数据库
2)当redis缓存的url小于定义的阀值
(二)并发同步
为了方便开发将模型设计为单生产者(master节点)多消费者模型(worker节点)。由于多了一层redis缓存的操作，与简单的单生产者和多消费者又有点不同。
这个不算简单的逻辑可以用如魔法一般的golang语言简单的实现(这段代码我是跪着写完的，哈哈哈)。带缓存的channel和阻塞式操作的redis。多说便无趣了，实现就在代码中。
```
##### 分发url实现
```
关键技术点:(1)与worker节点交互;(2)分发url时机
(一)与worker节点交互:
RPC、RPC、RPC重要的事情说三遍,单机的话可以基于Unix Socket协议;多机的话TCP无疑了
(二)合适的时机:
1)worker几点注册
2)worker工作完成
分别使用channel实现，贴下主要的代码:
for {
  select {
  case workAddr := <-m.regChan:
    work := &WorkInfo{workAddr : workAddr}
    m.workers[work] = true;
    fmt.Println("Register worker: ", work.workAddr)
    go dispatchJob(work, m)
  case workAddr := <-m.workDownChan:
    work := &WorkInfo{workAddr : workAddr}
    m.workers[work] = true;
    fmt.Println("WorkDown worker: ", work.workAddr)
    go dispatchJob(work, m)
  }
}
```

#### worker
```
完成master节点分发的任务，与爬虫模块交互。准确说应该是调用爬虫模块。实现注册和workdone接口
```
### Discussing
- [submit issue](https://github.com/zjucx/DistributedCrawler/issues/new)
- email: 862575451@qq.com
