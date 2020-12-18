package main

import (
	"myApp/internal/service"
)

func main() {
	s := service.New()

	s.Start()
}
