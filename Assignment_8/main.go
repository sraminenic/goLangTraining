package main

import (
	"webservice/routers"
)

func main() {
	r := routers.SetupRouter()
	r.Run(":8080")
}


