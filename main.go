package main

import "github.com/jpeizer/placehold/api/routers"

func main() {
	r := routers.SetupRouter()

	err := r.Run(":80")
	if err != nil {
		panic(err)
	}
}
