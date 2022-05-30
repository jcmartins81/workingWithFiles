package main

import (
	"log"
	"os"
)

func main() {
	err := os.Truncate("test.txt", 0)
	if err != nil {
		log.Fatal(err)
	}
}
