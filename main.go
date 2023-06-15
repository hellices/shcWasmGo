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

		if url == "/" {
			url = url + "index"
		}

		switch {
		case strings.Contains(url, "web.wasm"):
			w.Header().Set("Content-Type", "application/wasm")
			resp, _ := spinhttp.Get("https://raw.githubusercontent.com/hellices/shcWasmGo/poc/static/wasm/web.wasm")
			data, _ = ioutil.ReadAll(resp.Body)
			w.Write([]byte(data))

		case strings.Contains(url, "dashboard.js"):
			w.Header().Set("Content-Type", "application/javascript")
			// get from raw.githubusercontent
			resp, _ := spinhttp.Get("https://raw.githubusercontent.com/hellices/shcWasmGo/poc/static/js/dashboard.js")
			data, _ = ioutil.ReadAll(resp.Body)
			w.Write([]byte(data))

		case strings.Contains(url, "wasm_exec.js"):
			w.Header().Set("Content-Type", "application/javascript")
			// get from raw.githubusercontent
			resp, _ := spinhttp.Get("https://raw.githubusercontent.com/hellices/shcWasmGo/poc/static/js/wasm_exec.js")
			data, _ = ioutil.ReadAll(resp.Body)
			w.Write([]byte(data))

		case strings.Contains(url, "/demo/spin"):
			w.Header().Set("Content-Type", "text/html")
			resp, _ := spinhttp.Get("https://raw.githubusercontent.com/hellices/shcWasmGo/poc/static/html/demo_spin.html")
			data, _ = ioutil.ReadAll(resp.Body)
			w.Write([]byte(data))

		case strings.Contains(url, "/index"):
			w.Header().Set("Content-Type", "text/html")
			resp, _ := spinhttp.Get("https://raw.githubusercontent.com/hellices/shcWasmGo/poc/static/html/index.html")
			data, _ = ioutil.ReadAll(resp.Body)
			w.Write([]byte(data))

		case strings.Contains(url, "/runtime"):
			w.Header().Set("Content-Type", "text/html")
			resp, _ := spinhttp.Get("https://raw.githubusercontent.com/hellices/shcWasmGo/poc/static/html/runtime.html")
			data, _ = ioutil.ReadAll(resp.Body)
			w.Write([]byte(data))

		default:
			w.WriteHeader(404)
			w.Header().Set("Content-Type", "text/html")
			resp, _ := spinhttp.Get("https://raw.githubusercontent.com/hellices/shcWasmGo/poc/static/html/404.html")
			data, _ = ioutil.ReadAll(resp.Body)
			w.Write([]byte(data))
		}
	})
}

func main() {}
