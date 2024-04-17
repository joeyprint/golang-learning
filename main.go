package main

import (
	"fmt"
	"live/authentication/src/routers"
)

func main() {
	fmt.Println("Hello World!")
	routers.Router().Run(":8080")
}
