package main

import (
	"fmt"

	"github.com/jpeizer/placehold/api/routers"
)

func main() {
	fmt.Println("Starting server...")
	r := routers.SetupRouter()

	fmt.Println("Listening on port 80...")
	err := r.Run(":80")
	if err != nil {
		panic(err)
	}
}
