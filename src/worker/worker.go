package redismq

import (
  "net"
  "net/rpc"
)
type Worker {
  l net.Listener
  WorkerAddress string
}

func InitWorker(WorkerAddress string, nRPC int) *Worker{
  w := &Worker{}
  w.nRPC = nRPC
  w.WorkerAddress = WorkerAddress
}

func (w *Worker) StartRpcServer() {
  rpcs := rpc.NewServer()
  rpcs.Register(w)
  l, e := net.Listen("unix", w.WorkerAddress)
  if e != nil {
		log.Fatal("RunWorker: worker ", me, " error: ", e)
	}
	wk.l = l
  //add your code here

  for wk.nRPC != 0 {
		conn, err := wk.l.Accept()
		if err == nil {
			wk.nRPC -= 1
			go rpcs.ServeConn(conn)
			wk.nJobs += 1
		} else {
			break
		}
	}
	wk.l.Close()
	DPrintf("RunWorker %s exit\n", me)
}
