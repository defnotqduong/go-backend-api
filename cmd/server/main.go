package main

import "go/go-backend-api/internal/initialize"

func main() {
	// r := routers.NewRouter()
	// initmySQl()
	// initRedis()
	// initKafka()

	// r.Run(":8000")

	initialize.Run()
}
