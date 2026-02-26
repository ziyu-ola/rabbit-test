package main

import (
	"fmt"

	"github.com/ziyu-ola/rabbit-test/services"
)

func main() {
	g := services.NewGreeter("World")
	fmt.Println(g.Greet())
}
