package main

import (
	"fmt"
	"github.com/go-react/server/render"
	"net/http"
)

func main() {
	render.Init()
	fmt.Printf("Server is Listen on port 30000, happy hacking!\n")
	_ = http.ListenAndServe("0.0.0.0:30000", nil)
}
