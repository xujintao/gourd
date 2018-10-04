package main

import (
	"github.com/xujintao/gorge/http"
)

func main() {

	r := http.SetupRouter()
	r.Run() // listen and serve on 0.0.0.0:8080
}
