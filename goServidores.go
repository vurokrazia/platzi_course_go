package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	home := time.Now()
	canal := make(chan string)
	servidores := []string{
		"http://fb.com",
		"http://google.com",
		"http://facebook.com",
		"http://instagram.com",
	}

	count := 1
	for {
		for _, servidor := range servidores {
			go reviewServidor(servidor, canal)
		}
		//time.Sleep(1 * time.Second)
		fmt.Println(<-canal)
		if count == 100000 {
			break
		}
		count++
		fmt.Printf("counter: %s", string(count))
	}

	timeLast := time.Since(home)
	fmt.Printf("Time ejecution: %s", timeLast)
}

func reviewServidor(servidor string, canal chan string) {
	_, err := http.Get(servidor)
	if err != nil {
		//fmt.Println(servidor, "dont found :C")
		canal <- servidor + "dont found :C"
	} else {
		//fmt.Println(servidor, "Success!")
		canal <- servidor + "Success!"
	}
}
