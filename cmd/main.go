package main

import (
	"fmt"
	"live/authentication/routers"
)

func main() {
	fmt.Println("Hello World!")
	routers.Router().Run(":8080")
}
