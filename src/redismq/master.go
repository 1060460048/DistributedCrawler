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
}
"root:1234567890@/test?charset=utf8"

func InitMaster(dbname, dbstring, MasterAddress string) *Master {
  m = &Master{}
  m.dbname = dbname
  m.dbstring = dbstring
  m.MasterAddress = MasterAddress
  mr.alive = true
  return m
}

func RunMaster(dbname, dbstring string) {
  m := InitMaster(dbname, dbstring)
	l := sync.Mutex
  m.StartRpcServer()


	/*get_url_from_mysql := func(m *Master) string {
		var res = -1
		l.Lock()
		defer l.Unlock()
    rows, err := m.QueryUrls("select * from urls")
    if err != nil {
      fmt.Println("open mysql err:" + err);
    }
    for rows.Next() {
       var url string
       err = rows.Scan(&url)
       if err != nil {
         fmt.Println("open mysql err:" + err);
       }
       fmt.Println(url)
    }
	}

	add_url_to_redis := func(m *Master) {
		l.Lock()
		defer l.Unlock()
	}

  get_url_to_redis := func(m *Master) {
		l.Lock()
		defer l.Unlock()
	}*/
}

func (m *Master) StartRpcServer() {
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

func (m *Master) QueryUrls(sql string) (rows, err){
  db, err := sql.Open(m.dbname, m.dbstring)
  defer db.Close()
  if err != nil {
    fmt.Println("open mysql err:" + err);
  }
  rows, err := db.Query(sql)
}

func (m *Master) RedisAdd(){
  rs, err := redis.Dial("tcp", "127.0.0.1:6379")
  defer rs.Close()
  if err != nil {
    fmt.Println("Redis connection err: " + err)
  }
}
