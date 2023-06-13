package main

import (
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/hello", func(w http.ResponseWriter, req *http.Request) {
		w.Write([]byte("Hello World"))
	})

	http.HandleFunc("/robots", func(w http.ResponseWriter, req *http.Request) {
		// GET 호출
		resp, err := http.Get("https://google.com/robots.txt")

		if err != nil {
			panic(err)
		}

		defer resp.Body.Close()

		// 결과 출력
		data, err := ioutil.ReadAll(resp.Body)

		if err != nil {
			panic(err)
		}

		w.Header().Set("Content-Type", "text/plain")
		w.Write([]byte(data))
	})

	log.Panic(http.ListenAndServe(":50001", nil))

}
