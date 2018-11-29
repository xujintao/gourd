package main

import (
	"github.com/xujintao/gorge/apps/tpl/http"
)

func main() {

	r := http.SetupRouter()
	r.Run() // listen and serve on 0.0.0.0:8080
}
