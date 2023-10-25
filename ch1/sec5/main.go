package main

import (
	"fmt"
	"io"
	"net/http"
)

func fetch(url string) string {
	resp, err := http.Get(url)
	if err == nil {
		b, err := io.ReadAll(resp.Body)
		_ = resp.Body.Close()
		if err == nil {
			return string(b)
		}
	}
	return ""
}

func main() {
	url := "http://gopl.io"
	fmt.Println(fetch(url))
}
