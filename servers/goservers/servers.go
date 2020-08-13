package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	initTime := time.Now()
	channel := make(chan string)

	servers := []string{
		"http://google.com",
		"http://facebook.com",
		"http://wikipedia.com",
	}

	for _, server := range servers {
		go checkServers(server, channel)
		// fmt.Println("server:", server)
	}

	for i := 0; i < len(servers); i++ {
		fmt.Println(<-channel)
	}

	allTime := time.Since(initTime)
	fmt.Printf("time: %s", allTime)

	// time.Sleep(5 * time.Second)
	// fmt.Println("Done!")

}

func checkServers(server string, channel chan string) {
	fmt.Println("connecting to server: .....", server)
	_, err := http.Get(server)
	if err != nil {
		// fmt.Println(server, " is not available")
		channel <- server + " is not available"
	} else {
		// fmt.Println(server, " is working")
		channel <- server + " is working"
	}
}
