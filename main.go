package main

import (
	"net/http"

	"strings"

	spinhttp "github.com/fermyon/spin/sdk/go/http"
	"github.com/wasm_poc/index"
)

func init() {

	spinhttp.Handle(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/html")

		if strings.Contains(r.URL.RequestURI(), "poc") {
			data := index.GetWasm()
			w.Write([]byte(data))
		} else {
			data := index.GetIndex()
			w.Write([]byte(data))
		}
	})
}

func main() {}
