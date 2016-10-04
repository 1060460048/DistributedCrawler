package redismq

import (
  "net"
  "net/rpc"
  "fmt"
  //"container/list"
  //"github.com/mediocregopher/radix.v2/redis"
)

type Worker struct {
  l net.Listener
  nRPC int
  nJobs int
  WorkerAddress string
}

func initWorker(WorkerAddress string, nRPC int) *Worker{
  w := new(Worker)
  w.nRPC = nRPC
  w.WorkerAddress = WorkerAddress
  return w
}

func RunWorker(masterAddress, workerAddress string, nRPC int) {
  w := initWorker(workerAddress, nRPC)
  rpcs := rpc.NewServer()
  rpcs.Register(w)
  l, e := net.Listen("unix", w.WorkerAddress)
  if e != nil {
		fmt.Println("RunWorker: worker ", w.WorkerAddress, " error: ", e)
	}
	w.l = l
  //add your code here
  register(masterAddress, w.WorkerAddress)
  for w.nRPC != 0 {
		conn, err := w.l.Accept()
		if err == nil {
			w.nRPC -= 1
			go rpcs.ServeConn(conn)
			w.nJobs += 1
		} else {
			break
		}
	}
	w.l.Close()
}

func register(masterAddress, workerAddress string) {
  args := &RegisterArgs{}
  args.Worker = workerAddress
  var reply RegisterReply
  call(masterAddress, "Master.Register", args, &reply)
}

func (w *Worker) Dojob(args *DojobArgs, res *DojobReply) error {
  fmt.Println("DoJob: JobType ", args.JobType)
  switch args.JobType {
  case "Crawl":
    //DoCrawl(DojobArgs.Url)
    //DoAddRedis(args.Url)
  }
  return nil
}

/*func DoAddRedis(l list){
  conn, err := redis.Dial("tcp", "127.0.0.1:6379")
  defer conn.Close()
  if err != nil {
    fmt.Println("Redis connection err: %s", err)
  }
  for url := l.Front; url != nil; url = url.Next() {
    resp := conn.Cmd("RPUSH", "url", url)
    if resp.Err != nil {
      fmt.Println("Redis resp err: %s",err)
    }
  }
}*/
