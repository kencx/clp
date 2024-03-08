package main

import (
	"log"
	"os"
)

func main() {
	accessFile := "./access.log"
	f, err := os.Open(accessFile)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	results, err := parseFile(f)
	if err != nil {
		log.Fatal(err)
	}

	values, err := CountUserAgents(results)
	if err != nil {
		log.Fatal(err)
	}
	values.PrintTopN(10)
}
