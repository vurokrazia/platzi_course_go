package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	home := time.Now()
	servidores := []string{
		"http://fb.com",
		"http://google.com",
		"http://facebook.com",
		"http://instagram.com",
	}
	for _, servidor := range servidores {
		reviewServidor(servidor)
	}
	timeLast := time.Since(home)
	fmt.Printf("Time ejecution: %s", timeLast)
}

func reviewServidor(servidor string) {
	_, err := http.Get(servidor)
	if err != nil {
		fmt.Println(servidor, "dont found :C")
	} else {
		fmt.Println(servidor, "Success!")
	}
}
