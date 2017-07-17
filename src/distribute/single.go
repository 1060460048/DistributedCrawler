package distribute

import (
  "fmt"
  "time"
  // "errors"
  "scrawler"
  "model"
)

type Single struct {
  ThreadNum           int
  jobChan             chan string
  pool                ThreadPool
  rmq                 *model.RedisMq
}

func initSingle(threadNum int, jobNum int) (s *Single) {
  rmq, err := model.InitRedisMq("127.0.0.1:6379", 1)
  if err != nil {
    fmt.Println(time.Now().Format("2006-01-02 15:04:05") + " single.go initSingle error" + err.Error())
    return nil
  }
  s = &Single{}
  s.jobChan = make(chan string, jobNum)
  s.pool.Init(threadNum, jobNum)
  go s.pool.Start()
  s.rmq = rmq
  fmt.Println(time.Now().Format("2006-01-02 15:04:05") + " single.go initSingle success")
  return s
}

func RunSingle(threadNum int, jobNum int, startUrl string) {
  fmt.Println(time.Now().Format("2006-01-02 15:04:05") + " single.go RunSingle start")

  s := initSingle(threadNum, jobNum)

  go s.rmq.LoadUrlsFromRedis(s.jobChan)

  scrawler.Scrawler(startUrl)

  for {
    url,ok := <-s.jobChan
    if !ok {
        break;
    }
    fmt.Println(time.Now().Format("2006-01-02 15:04:05") + " single.go RunSingle add task")
    s.pool.AddTask(func () error {
      return scrawler.Scrawler(url)
    })
  }
  s.pool.Stop()
  s.rmq.Mgo.Session.Close()
  s.rmq.C.Close()
}
