package redismq

import (
  "fmt"
  //"sync"
  "net"
  "net/rpc"
  //"github.com/mediocregopher/radix.v2/redis"
  //"github.com/garyburd/redigo/redis"
)

type Master struct {
  dbname string
  //dbstring string
  MasterAddress string
  l               net.Listener
  alive           bool
  registerChannel chan string
  workDownChnanel chan string
  urlChannel chan string
  workers map[WorkInfo]bool
}

type WorkInfo struct {
  workAddr string
}

func initMaster(dbname, MasterAddress string) *Master {
  m := &Master{}
  m.dbname = dbname
  m.MasterAddress = MasterAddress
  m.alive = true
  m.registerChannel = make(chan string)
  m.urlChannel = make(chan string)
  m.workers = make(map[WorkInfo]bool)
  return m
}

func RunMaster(dbname, mr string) {
  m := initMaster(dbname, mr)
  go startRpcServer(m)
  fmt.Println("Master has run: ", mr)
  for {
    select {
    case workAddr := <-m.registerChannel:
      work := WorkInfo{workAddr : workAddr}
      m.workers[work] = true;
      fmt.Println("Register worker: ", work.workAddr)
      DispatchJob(work)
    case workAddr := <-m.workDownChnanel:
      work := WorkInfo{workAddr : workAddr}
      m.workers[work] = true;
      fmt.Println("WorkDown worker: ", work.workAddr)
      DispatchJob(work)
    }
  }
}

func DispatchJob(workInfo WorkInfo) {
  /*conn, err := redis.Dial("tcp", "127.0.0.1: ")
  defer conn.Close()
  if err != nil {
    fmt.Println("Redis connection err: %s", err)
  }

  url, err := conn.Cmd("BLPOP", "url", 0).Str()
  if err != nil {
    fmt.Println(err)
  }
  fmt.Println("DispatchJob: ", url)*/

  args := new(DojobArgs)
  args.Url = "www.baidu.com"//url
  args.JobType = "hehe"
  var reply DojobReply
  call(workInfo.workAddr, "Worker.Dojob", args, &reply)
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
  rpcs := rpc.NewServer()
  rpcs.Register(m)
  l, e := net.Listen("unix", m.MasterAddress)
  if e != nil {
		fmt.Println("RegstrationServer", m.MasterAddress, " error: ", e)
	}
	m.l = l
  // now that we are listening on the master address, can fork off
	// accepting connections to another thread.
	go func() {
		for m.alive {
			conn, err := m.l.Accept()
			if err == nil {
				go rpcs.ServeConn(conn)
			} else {
				fmt.Println("RegistrationServer: accept error", err)
				break
			}
      conn.Close()
		}
		fmt.Println("RegistrationServer: done\n")
	}()
}

/*func DispatchUrl(m *Master) {
  conn, err := redis.Dial("tcp", "127.0.0.1:6379")
  defer conn.Close()
  if err != nil {
    fmt.Println("Redis connection err: %s", err)
  }

  go func (){
    for {
      // resp := conn.Cmd("HMSET", "album:1", "title", "Electric Ladyland", "artist", "Jimi Hendrix", "price", 4.95, "likes", 8)
      // // Check the Err field of the *Resp object for any errors.
      // if resp.Err != nil {
      //     log.Fatal(resp.Err)
      // }
      //
      // fmt.Println("Electric Ladyland added!")
      select{
      case workerAddr:= <-m.workDownChnanel:
        url, err := conn.Cmd("BLPOP", "url", 0).Str()
        if err != nil {
          fmt.Println(err)
        }
        args := new(DojobArgs)
        args.Url = url
        var reply DojobReply
        ok := call(workerAddr, "Worker.Dojob", args, &reply)
      }
    }
  }()

}*/
