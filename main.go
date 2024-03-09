package main

import (
	"fmt"
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

	results, err := decodeFile(f)
	if err != nil {
		log.Fatal(err)
	}

	values, err := CountRemoteIPs(results)
	if err != nil {
		log.Fatal(err)
	}
	values.PrintTopN(10)

	uv, err := UniqueVisitors(results)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(uv)

	pv, err := PageViews(results)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(pv)
}
