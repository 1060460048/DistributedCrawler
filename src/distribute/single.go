package distribute

import (
  "fmt"
  "scrawler"
  "model"
)

type Single struct {
  ThreadNum           int
  jobChan             chan string
  pool                ThreadPool
  rmq                 *model.RedisMq
}

func initSingle(threadNum int, jobNum int) s *Single {
  s = &Single{}
  s.JobChan = make(chan string, jobNum)
  s.pool.Init(threadNum, jobNum)
  s.pool.Start()
  s.rmq, _ = model.InitRedisMq("127.0.0.1:32770", 1)
}

func RunSingle(threadNum int, jobNum int, startUrl string) {
  fmt.Println("=======RunMaster Begin=======")

  s = initSingle(threadNum, jobNum)

  go loadUrlsFromRedis(s)

  scrawler(startUrl)

  for {
    url,ok := <-s.jobChan
    if !ok {
        break;
    }
    s.pool.AddTask(func () error {
      return scrawler(url)
    })
  }
  s.pool.Stop()
  s.rmq.C.Close()
}

func scrawler(url string) error{
  scrawler(url)
  return nil
}
