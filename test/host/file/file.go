package main

import (
	"io/ioutil"
	"log"
)

func main() {
	data, err := ioutil.ReadFile("./Text")

	if err != nil {
		log.Panic(err)
	}

	log.Println("File Contents : ", string(data))
}