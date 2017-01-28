package distribute

import (
  "net/rpc"
  "net/http"
  "fmt"
)

type Worker struct {
  addr string
  addUrlChannel chan bool
}

func initWorker(addr string) *Worker{
  w := &Worker{}
  w.addr = addr
  w.addUrlChannel = make(chan bool)
  return w
}

func RunWorker(mAddr, wAddr string) {
  fmt.Println("=======RunWorker Begin=======")
  w := initWorker(wAddr)
  // go startRpcWorker(w)
  register(mAddr, w.addr)
  fmt.Println("=======RunWorker End=======")
  // for {
  //   select {
  //   case: <- addUrlChannel
  //     addUrlsToMongodb()
  //   }
  // }
  // defer w.mgo.MgoClient.Close()
}

func register(mAddr, wAddr string) {
  args := &RegisterArgs{}
  args.Worker = wAddr
  var reply RegisterReply
  call(mAddr, "Master.Register", args, &reply)
}

func (w *Worker) Dojob(args *DojobArgs, res *DojobReply) error {
  fmt.Println("DoJob: JobType ", args.JobType)
  switch args.JobType {
  case "Crawl":
    //save your pages and return the next urls
    //urls := scrawler.DoCrawl(DojobArgs.Urls)
    /*
    if urls.length() > 0 {
      addUrlChannel <- true
    }
     */
  }
  return nil
}

func startRpcWorker(w *Worker) {
  //need code reconstruction
  rpc.Register(w)
  rpc.HandleHTTP()
  err := http.ListenAndServe(w.addr, nil)
  fmt.Println("RegistrationServer: accept error", err)
  // rpcs := rpc.NewServer()
  // rpcs.Register(w)
  // l, e := net.Listen("tcp", w.addr)
  // if e != nil {
	// 	fmt.Println("RunWorker: worker ", w.addr, " error: ", e)
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
