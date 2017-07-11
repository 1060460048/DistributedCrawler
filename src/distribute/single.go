package distribute

import (
  "fmt"
  "scrawler"
)

type Single struct {
  ThreadNum           int
  Jobs                chan func() error
  Result              chan error
  jobChan             chan string
  pool                ThreadPool
}

func initSingle(threadNum int, jobNum int) s *Single {
  s = &Single{}
  s.ThreadNum = threadNum
  s.Jobs = make(chan func() error, threadNum)
  s.Result = make(chan error, threadNum)
  s.pool.Init(threadNum, jobNum)
  s.pool.Start()
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
}

func scrawler(url string) error{
  scrawler(url)
  return nil
}
