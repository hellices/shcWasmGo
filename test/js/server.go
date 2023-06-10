package main

import (
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	content, err := ioutil.ReadFile("./main.wasm")

	if err != nil {
		log.Panic(err)
	}

	http.HandleFunc("/main.wasm", func(w http.ResponseWriter, r *http.Request) {
		if err != nil {
			w.WriteHeader(404)
			w.Write([]byte(http.StatusText(404)))
			return
		}
		// Allow CORS
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
		w.Header().Add("Content-Type", "application/wasm")
		w.Write(content)
	})

	log.Fatal(http.ListenAndServe(":50001", nil))
}
