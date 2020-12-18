package main

import (
	"flag"
	"log"
	"myApp/internal/service"
)

var addr string

func init() {
	flag.StringVar(&addr, "addr", ":0", "")
	flag.Parse()
}

func main() {
	s := service.New()
	if err != nil {
		log.Fatalf("err: %v", err)
	}
	s.Start()
}
