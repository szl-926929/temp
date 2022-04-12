package main

import (
	"time"

	"github.com/go-resty/resty/v2"
)

func main() {
	client := resty.New()
	client.SetTimeout(5 * time.Second).R().Get("http://localhost:8080/api")
}
