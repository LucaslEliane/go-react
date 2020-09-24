package main

import "net/http"

func main() {
	_ = http.ListenAndServe("0.0.0.0:30000", nil)
}
