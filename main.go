package main

import (
	"fmt"
	router "live/authentication/src"
)

func main() {
	fmt.Println("Hello World!")
	router := router.Router()
	router.Run(":8080")
}
