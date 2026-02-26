package main

import (
	"fmt"
	"os"

	"github.com/ziyu-ola/rabbit-test/services"
)

func main() {
	g := services.NewGreeter("World")
	fmt.Println(g.Greet())

	if len(os.Args) > 1 {
		birthday := os.Args[1]
		age, err := services.AgeFromBirthdayString(birthday)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error: %v\n", err)
			os.Exit(1)
		}
		fmt.Printf("Age: %d\n", age)
	}
}
