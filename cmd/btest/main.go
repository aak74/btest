package main

import (
	"log"
	"net/http"
)

func main() {
	err := http.ListenAndServe(":8090", NewServer())
	if err != nil {
		log.Fatal(err)
	}
}
