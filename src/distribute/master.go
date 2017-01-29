package distribute

import (
  "fmt"
  //"sync"
  // "net"
  "net/rpc"
  "net/http"
  //"github.com/mediocregopher/radix.v2/redis"
  // "github.com/garyburd/redigo/redis"
  "model"
  // "time"
)

type Master struct {
  addr             string
  // l                 net.Listener
  // alive            bool
  regChan             chan string
  workDownChan        chan string
  urlChan             chan string
  workers             map[*WorkInfo]bool
  rmq                 *model.RedisMq
}

type WorkInfo struct {
  workAddr string
}

func initMaster(addr string) (m *Master, err error) {
  m = &Master{}
  m.addr = addr
  // m.alive = true
  m.regChan = make(chan string)
  m.urlChan = make(chan string, 5)
  m.workers = make(map[*WorkInfo]bool)
  m.rmq, err = model.InitRedisMq("127.0.0.1:32770", 1)
  return m, err
}

func RunMaster(addr string) {
  fmt.Println("=======RunMaster Begin=======")
  m, err := initMaster(addr)
  if err != nil {
    fmt.Println("initMaster error: " + err.Error())
    return
  }
  defer m.rmq.C.Close()

  go startRpcMaster(m)
  go loadUrlsFromRedis(m)
  // go RunRedisMq(dbname, 0)
  fmt.Println("=======RunMaster End=======")
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
}

/*
 * this function is likely a consumer.
 */
func dispatchJob(workInfo *WorkInfo, m *Master) {
  var urls []string
  for i:= 0;i < 10;i++ {
    url := <- m.urlChan // get ulr from channel
    urls = append(urls, url)
  }
  args := &DojobArgs{}
  // args.Url = "www.baidu.com"//url
  args.JobType = "Crawl"
  args.Urls = urls
  var reply DojobReply
  err := call(workInfo.workAddr, "Worker.Dojob", args, &reply)
  if err == true {
    m.workers[workInfo] = false;
    fmt.Println("dispatchJob success worker: " + workInfo.workAddr)
  }
}

/*
 * this function is likely a producter.
 */
func loadUrlsFromRedis(m *Master) {
  //1) load Data
  //2) dispatchjob
  // When finish you need dispatchjob for
  // every blocked work because of none data in redis
  fmt.Println("loadUrlsFromRedis: begin")
  for {
    url := m.rmq.GetUrlBlock()
    // if len(urls) == 0 {
    //   fmt.Println("loadUrlsFromRedis urls is nil sleep 60s")
    //   time.Sleep(60 * time.Second)
    // }
    m.urlChan <- url
  }
}

func (m *Master) Register(args *RegisterArgs, res *RegisterReply) error {
   m.regChan <- args.Worker
   return nil
}

/*func (m *Master) JobFinish(args *FinishArgs, res *FinishReply) {
   fmt.Println("Finish worker:%s\n", args.Worker)
   m.workDownChan <- args.Worker
}*/

func startRpcMaster(m *Master) {
  rpc.Register(m)
  rpc.HandleHTTP()
  err := http.ListenAndServe(m.addr, nil)
  if err != nil {
    fmt.Println("RegistrationServer: accept error", err)
  }
  fmt.Println("startRpcMaster: success")
}
