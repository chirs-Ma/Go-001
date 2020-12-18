package main

import (
	"log"
	"myApp/internal/service"
)

func main() {
	s := service.New()
	if err != nil {
		log.Fatalf("err: %v", err)
	}
	s.Start()
}
