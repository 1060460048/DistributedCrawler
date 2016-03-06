package redismq

import (
  "fmt"
  "database/sql"
  "sync"
  _ "github.com/go-sql-driver/mysql"
  "github.com/garyburd/redigo/redis"
)

type Master struct{
  dbname string
  dbstring string

}
"root:1234567890@/test?charset=utf8"

func InitMaster(dbname, dbstring string) *Master {
  m = &Master{}
  m.dbname = dbname
  m.dbstring = dbstring
  return m
}

func Run(dbname, dbstring string) {
  m := InitMaster(dbname, dbstring)
	l := sync.Mutex

	get_url_from_mysql := func(m *Master) string {
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
	}
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
