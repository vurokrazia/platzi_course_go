package main

import (
	"fmt"
	"io"
	"net/http"
)

type desktopWeb struct {
}

func (desktopWeb) Write(p []byte) (int, error) {
	fmt.Println(string(p))
	return len(p), nil
}

func main() {
	resp, err := http.Get("https://fb.com")
	if err != nil {
		fmt.Println(err)
	}
	e := desktopWeb{}
	io.Copy(e, resp.Body)
}
