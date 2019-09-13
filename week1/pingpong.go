package main

import "fmt"

func pinger(ping chan int, pong chan int) {
	for {
		<-ping
		fmt.Println("Ping")
		pong <- 1
	}
}

func ponger(ping chan int, pong chan int) {
	for {
		<-pong
		fmt.Println("Pong")
		ping <- 1
	}
}

func mainPingpong() {
	ping := make(chan int)
	pong := make(chan int)
	// ping <- 1
	go pinger(ping, pong)
	go ponger(ping, pong)
	ping <- 1
	fmt.Scanln()
}
