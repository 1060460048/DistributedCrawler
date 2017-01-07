package redismq

import (
  "net"
  "net/rpc"
  "fmt"
  "dscrawl"
  //"container/list"
  //"github.com/mediocregopher/radix.v2/redis"
)

type Worker struct {
  l net.Listener
  nRPC int
  nJobs int
  Address string
  mgosess *mgo.Session
}

func initWorker(Address string, nRPC int) *Worker{
  w := new(Worker)
  w.nRPC = nRPC
  w.Address = Address
  w.mgosess = dscrawl.InitDB
  return w
}

func RunWorker(masterAddress, workerAddress string, nRPC int) {
  w := initWorker(workerAddress, nRPC)
  go startRpcServer(w)
  register(masterAddress, w.Address)
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
    //urls := scrawler.DoCrawl(DojobArgs.Url)
    //addUrlsToMongodb(w, urls)
  }
  return nil
}

func addUrlsToMongodb(w *Worker, urls []string){

  defer w.mgo.MgoClient.CLose()
}

func startRpcServer(w *Worker) {
  //need code reconstruction
  rpc.Register(w)
  rpc.HandleHTTP()
  err := http.ListenAndServe(w.Address, nil)
  fmt.Println("RegistrationServer: accept error", err)
  // rpcs := rpc.NewServer()
  // rpcs.Register(w)
  // l, e := net.Listen("tcp", w.Address)
  // if e != nil {
	// 	fmt.Println("RunWorker: worker ", w.Address, " error: ", e)
	// }
	// w.l = l
  //add your code here
  // for w.nRPC != 0 {
	// 	conn, err := w.l.Accept()
	// 	if err == nil {
	// 		w.nRPC -= 1
	// 		go rpcs.ServeConn(conn)
	// 		w.nJobs += 1
	// 	} else {
	// 		break
	// 	}
	// }
	// w.l.Close()
}

/*func getUrlsFromRedis() (l list){
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
  return l
}*/
