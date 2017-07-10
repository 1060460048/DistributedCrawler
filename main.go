package main

import (
	"fmt"
	//"net/http"
	//"log"
	"os"
	"distribute"
	// "scrawler"
	// "segment"
)
// Can be run in 3 ways:
// 1) Master (e.g., go run main.go master localhost:7777)
// 2) Worker (e.g., go run main.go worker localhost:7777 localhost:7778 &)
// 3) Master or Worker (e.g., go run main.go worker sequential)
func main() {
  if len(os.Args) != 3 {
		//scrawler.Scrawler()
		distribute.RunWorker(os.Args[2], os.Args[3])
    fmt.Printf("%s: see usage comments in file\n", os.Args[0])
  } else if os.Args[1] == "master" {
    distribute.RunMaster(os.Args[2])
  } else if os.Args[1] == "worker" {
    distribute.RunWorker(os.Args[2], os.Args[3])
  } else if os.Args[1] == "single" {
		//线程数
		distribute.RunSingle(5)
    //redismq.RunSingle(5, 3, os.Args[2], Map, Reduce)
	} else {
		//fmt.Printf("%s: see usage comments in file\n", os.Args[0])
	}
}

/*type StatisticsServer struct {
	port     string
}

func InitServer(port string) *StatisticsServer {
  s = &StatisticsServer{
		port:    port,
	}
  return s
}

func (s *StatisticsServer) Start() {
	go func () {
		log.Printf("STARTING REDISMQ SERVER ON PORT %s", s.port)
		err := http.ListenAndServe(":"+s.port, nil)
		if err != nil {
			log.Fatalf("REDISMQ SERVER SHUTTING DOWN [%s]\n\n", err.Error())
		}
	}()
}*/
