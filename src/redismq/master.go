package redismq

import (
  "fmt"
  "database/sql"
  "sync"
  "net"
  "net/rpc"
  _ "github.com/go-sql-driver/mysql"
  _ "github.com/garyburd/redigo/redis"
)

type Master struct{
  dbname string
  dbstring string
  MasterAddress string
  l               net.Listener
  alive           bool
  registerChannel chan string
}
"root:1234567890@/test?charset=utf8"

func InitMaster(dbname, dbstring, MasterAddress string) *Master {
  m = &Master{}
  m.dbname = dbname
  m.dbstring = dbstring
  m.MasterAddress = MasterAddress
  m.alive = true
  m.registerChannel = make(chan string)
  return m
}

func RunMaster(dbname, dbstring string) {
  m := InitMaster(dbname, dbstring)
	l := sync.Mutex
  go StartRpcServer(m)

  for {
    select {
    case workAddr := <-m.redisterChannel:
    }
  }
}

func (m *Master) Register(args *RegisterArgs, res *RegisterReply) {
   fmt.Println("Register worker:%s\n", args.Worker)
   m.registerChannel <- args.Worker
   res.Ok = true
   return nil
}

func StartRpcServer(m *Master) {
  rpcs := rpc.NewServer()
  rpcs.Register(m)
  l, e := net.Listen("unix", m.MasterAddress)
  if e != nil {
		log.Fatal("RegstrationServer", mr.MasterAddress, " error: ", e)
	}
	mr.l = l
  // now that we are listening on the master address, can fork off
	// accepting connections to another thread.
	go func() {
		for mr.alive {
			conn, err := mr.l.Accept()
			if err == nil {
				go func() {
					rpcs.ServeConn(conn)
					conn.Close()
				}()
			} else {
				log.Fatal("RegistrationServer: accept error", err)
				break
			}
		}
		log.Fatal("RegistrationServer: done\n")
	}()
}

func () DispatchUrl() {
  conn, err := redis.Dial("tcp", "127.0.0.1:6379")
  defer conn.Close()
  if err != nil {
    fmt.Println("Redis connection err: %s", err)
  }

  go func (){
    for {
      resp := conn.Cmd()
      if resp.Err != nil {
        fmt.Println("Redis resp err: %s",err)
      }
    }
  }()

}
