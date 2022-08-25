package main

import (
	"net/http"

	_ "net/http/pprof"

	_ "webconsole/generator"
)

func main() {
	http.ListenAndServe(":8001", nil)
}
