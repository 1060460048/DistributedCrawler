package redismq

import (
  "fmt"
  //"sync"
  "net"
  "net/rpc"
  "github.com/mediocregopher/radix.v2/redis"
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
  workers map[string]bool
}

func InitMaster(dbname, MasterAddress string) *Master {
  m := &Master{}
  m.dbname = dbname
  m.MasterAddress = MasterAddress
  m.alive = true
  m.registerChannel = make(chan string)
  m.urlChannel = make(chan string)
  m.workers = make(workers[string]bool)
  return m
}

func RunMaster(dbname, mr string) {
  m := InitMaster(dbname, mr)
  go StartRpcServer(m)

  for {
    select {
    case workAddr := <-m.registerChannel:
      m.workers[workAddr] = true;
      DispatchJob(workAddr)
    case workAddr :=m.workDownChnanel:
      m.workers[workAddr] = true;
      DispatchJob(workAddr)
    }
  }
}

func DispatchJob(url string) {
  args := new(DojobArgs)
  args.Url = url
  args.JobType = "hehe"
  var reply DojobReply
  ok := call(workerAddr, "Worker.Dojob", args, &reply)
}

func (m *Master) Register(args *RegisterArgs, res *RegisterReply) {
   fmt.Println("Register worker:%s\n", args.Worker)
   m.registerChannel <- args.Worker
}

func (m *Master) JobFinish(args *FinishArgs, res *FinishReply) {
   fmt.Println("Finish worker:%s\n", args.Worker)
   m.workDownChnanel <- args.Worker
}

func StartRpcServer(m *Master) {
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
      fmt.Println("RegstrationServer", m.MasterAddress, " m.alive ", m.alive)
			conn, err := m.l.Accept()
			if err == nil {
				go func() {
					rpcs.ServeConn(conn)
					conn.Close()
				}()
			} else {
				fmt.Println("RegistrationServer: accept error", err)
				break
			}
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
