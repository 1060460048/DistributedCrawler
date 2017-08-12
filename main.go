package main

import (
	"fmt"
	//"net/http"
	//"log"
	"os"
	"strconv"
	"distribute"
	// "scrawler"
	// "segment"
	"unsafe"
)
// Can be run in 3 ways:
// 1) Master (e.g., go run main.go master localhost:7777)
// 2) Worker (e.g., go run main.go worker localhost:7777 localhost:7778 &)
// 3) Master or Worker (e.g., go run main.go worker sequential)
func main() {
	if len(os.Args) < 2 {
		sz := 100
		p := make([]int, sz)
		p[2] = 5
		// INVALID: end points outside allocated space.
		// b := make([]byte, n)
		// end = unsafe.Pointer(uintptr(unsafe.Pointer(&b[0])) + uintptr(n))
		fmt.Printf("%s: see usage comments in file %d\n", os.Args[0], *(*int)(unsafe.Pointer(uintptr(unsafe.Pointer(&p[0])) + 2*unsafe.Sizeof(p[0]))))
		return
	}
	switch os.Args[1] {
	case "master":
		if len(os.Args) == 3 {
			distribute.RunMaster(os.Args[2])
		}
	case "worker":
		if len(os.Args) == 4 {
			distribute.RunWorker(os.Args[2], os.Args[3])
		}
	case "single":
		if len(os.Args) == 5 {
			threadNum, err1 := strconv.Atoi(os.Args[2])
			jobNum, err2 := strconv.Atoi(os.Args[3])
			if err1 != nil || err2 != nil {
				fmt.Printf("%s: see usage comments in file %d\n", os.Args[0], len(os.Args))
			}
			distribute.RunSingle(threadNum, jobNum, os.Args[4])
		} else {
			fmt.Printf("%s: see usage comments in file %d\n", os.Args[0], len(os.Args))
		}
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
