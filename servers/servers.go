package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	initTime := time.Now()

	servers := []string{
		"http://google.com",
		"http://facebook.com",
		"http://www.rubioharo.tech",
	}

	for _, servidor := range servers {
		checkServers(servidor)
	}

	allTime := time.Since(initTime)
	fmt.Printf("time: %s", allTime)
}

func checkServers(server string) {
	_, err := http.Get(server)
	if err != nil {
		fmt.Println(server, " is not available")
	} else {
		fmt.Println(server, " is working")
	}
}
