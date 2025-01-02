package main

import (
	"go/go-backend-api/internal/routers"
)

func main() {
	r := routers.NewRouter()
	r.Run(":8000")
}
