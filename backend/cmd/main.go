package main

import (
	"backend/delivery/http"
	"log"
)

func main() {

	http.StartServerHttp()

	if err := http.StartServerHttp(); err != nil {
		log.Fatalf("error %s", err)
	}
}
