package redismq

import (
  "fmt"
  //"sync"
  "net"
  "net/rpc"
  "net/http"
  //"github.com/mediocregopher/radix.v2/redis"
  //"github.com/garyburd/redigo/redis"
)

type Master struct {
  dbname string
  //dbstring string
  Address string
  // l               net.Listener
  // alive           bool
  registerChannel chan string
  workDownChnanel chan string
  readUrlsChannel chan bool
  urlChannel chan string
  workers map[WorkInfo]bool
  rmq *RedisMq
}

type WorkInfo struct {
  workAddr string
}

func initMaster(dbname, Address string) *Master {
  m := &Master{}
  m.dbname = dbname
  m.Address = Address
  // m.alive = true
  m.registerChannel = make(chan string)
  m.urlChannel = make(chan string, 100)
  m.workers = make(map[WorkInfo]bool)
  m.rmp = initRedisMq()
  return m
}

func RunMaster(dbname, mr string) {
  m := initMaster(dbname, mr)
  go startRpcServer(m)
  go loadUrlsFromRedis(m)
  // go RunRedisMq(dbname, 0)
  fmt.Println("Master has run: ", mr)
  for {
    select {
    case workAddr := <-m.registerChannel:
      work := WorkInfo{workAddr : workAddr}
      m.workers[work] = true;
      fmt.Println("Register worker: ", work.workAddr)
      go dispatchJob(work, m)
    case workAddr := <-m.workDownChnanel:
      work := WorkInfo{workAddr : workAddr}
      m.workers[work] = true;
      fmt.Println("WorkDown worker: ", work.workAddr)
      go dispatchJob(work, m)
    // case <-m.readUrlsChannel:
    //   loadUrlsToRedis(m)
    // }
  }
}

/*
 * this function is likely a consumer.
 */
func dispatchJob(workInfo WorkInfo, m Master) {
  urls := []
  for i:= 0;i < 10;i++ {
    url <- m.urlChannel // get ulr from channel
    urls = append(urls, url)
  }
  m.workers[workInfo] = false;
  args := new(DojobArgs)
  // args.Url = "www.baidu.com"//url
  args.JobType = "Crawl"
  args.Urls = urls
  var reply DojobReply
  call(workInfo.workAddr, "Worker.Dojob", args, &reply)
}

/*
 * this function is likely a producter.
 */
func loadUrlsFromRedis(m Master) {
  //1) load Data
  //2) dispatchjob
  // When finish you need dispatchjob for
  // every blocked work because of none data in redis
  for {
    urls := m.rmq.GetUrls()
    for i, v := range urls {
      m.urlChannel <- v
    }
  }
}

func (m *Master) Register(args *RegisterArgs, res *RegisterReply) error {
   m.registerChannel <- args.Worker
   return nil
}

/*func (m *Master) JobFinish(args *FinishArgs, res *FinishReply) {
   fmt.Println("Finish worker:%s\n", args.Worker)
   m.workDownChnanel <- args.Worker
}*/

func startRpcServer(m *Master) {
  rpc.Register(m)
  rpc.HandleHTTP()
  err := http.ListenAndServe(m.Address, nil)
  fmt.Println("RegistrationServer: accept error", err)
}
