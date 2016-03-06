package main

import (
	"fmt"
	"net/http"
	"log"
)

type StatisticsServer struct {
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
}
