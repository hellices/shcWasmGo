package main

import (
	"io/ioutil"
	"log"
)

func main() {
	data, err := ioutil.ReadFile("../static/404.html")

	if err != nil {
		log.Panic(err)
	}

	log.Print(string(data))
}
