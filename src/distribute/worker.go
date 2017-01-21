package distribute

import (
  "net"
  "net/rpc"
  "fmt"
)

type Worker struct {
  // l net.Listener
  // nRPC int
  // nJobs int
  Address string
  // mgosess *mgo.Session
  addUrlChannel chan bool
}

func initWorker(Address string, nRPC int) *Worker{
  w := &Worker{}
  // w.nRPC = nRPC
  w.Address = Address
  // w.mgosess = dscrawl.InitDB()
  w.addUrlChannel = make(chan bool)
  return w
}

func RunWorker(masterAddress, workerAddress string, nRPC int) {
  w := initWorker(workerAddress, nRPC)
  go startRpcServer(w)
  register(masterAddress, w.Address)
  // for {
  //   select {
  //   case: <- addUrlChannel
  //     addUrlsToMongodb()
  //   }
  // }
  // defer w.mgo.MgoClient.Close()
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
    //save your pages and return the next urls
    //urls := scrawler.DoCrawl(DojobArgs.Url)
    /*
    if urls.length() > 0 {
      addUrlChannel <- true
    }
     */
  }
  return nil
}

// func addUrlsToMongodb(w *Worker, urls []string){
//   w.mgo.InsertUrls(urls)
// }

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
