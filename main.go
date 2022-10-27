package main

import (
	"submission-3/routers"
)

func main() {
	var PORT = ":8080"
	routers.StartServer().Run(PORT)
}
