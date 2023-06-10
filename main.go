package main

import (
	"io/ioutil"
	"net/http"
	"strings"

	spinhttp "github.com/fermyon/spin/sdk/go/http"
)

func init() {

	spinhttp.Handle(func(w http.ResponseWriter, r *http.Request) {

		url := r.URL.RequestURI()
		var data []byte

		switch {
		case strings.Contains(url, "dashboard.js"):
			w.Header().Set("Content-Type", "application/javascript")
			// get from raw.githubusercontent
			resp, _ := spinhttp.Get("https://raw.githubusercontent.com/hellices/shcWasmGo/poc_test/static/dashboard.js")
			data, _ = ioutil.ReadAll(resp.Body)
			w.Write([]byte(data))
		case strings.Contains(url, "/demo/spin"):
			w.Header().Set("Content-Type", "text/html")
			resp, _ := spinhttp.Get("https://raw.githubusercontent.com/hellices/shcWasmGo/poc_test/static/demo_spin.html")
			data, _ = ioutil.ReadAll(resp.Body)
			w.Write([]byte(data))
		case strings.Contains(url, "/"):
			w.Header().Set("Content-Type", "text/html")
			resp, _ := spinhttp.Get("https://raw.githubusercontent.com/hellices/shcWasmGo/poc_test/static/index.html")
			data, _ = ioutil.ReadAll(resp.Body)
			w.Write([]byte(data))
		default:
			w.WriteHeader(404)
			w.Header().Set("Content-Type", "text/html")
			resp, _ := spinhttp.Get("https://raw.githubusercontent.com/hellices/shcWasmGo/poc_test/static/404.html")
			data, _ = ioutil.ReadAll(resp.Body)
			w.Write([]byte(data))
		}
	})
}

func main() {}
