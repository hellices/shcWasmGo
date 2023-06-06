package main

import (
	"net/http"
	"strings"

	spinhttp "github.com/fermyon/spin/sdk/go/http"
	"github.com/wasm_poc/index"
)

func init() {

	spinhttp.Handle(func(w http.ResponseWriter, r *http.Request) {

		if strings.Contains(r.URL.RequestURI(), ".js") {
			w.Header().Set("Content-Type", "application/javascript")
			data := index.GetDashboard()
			w.Write([]byte(data))
		} else {
			if strings.Contains(r.URL.RequestURI(), "poc") {
				w.Header().Set("Content-Type", "text/html")
				data := index.GetWasm()
				w.Write([]byte(data))
			} else {
				w.Header().Set("Content-Type", "text/html")
				data := index.GetIndex()
				w.Write([]byte(data))
			}
		}

	})
}

func main() {}
