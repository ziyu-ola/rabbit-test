package main

import (
	"fmt"
	"os"

	"github.com/ziyu-ola/rabbit-test/db"
	"github.com/ziyu-ola/rabbit-test/services"
)

func lookupUsers() {
	if err := db.InitDB(); err != nil {
		fmt.Fprintf(os.Stderr, "db init error: %v\n", err)
		return
	}
	for uid := 1000; uid <= 1015; uid++ {
		name, err := db.GetNameByUid(uid)
		if err != nil {
			fmt.Printf("uid %d: error: %v\n", uid, err)
			continue
		}
		fmt.Printf("uid %d: %s\n", uid, name)
	}
}

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

	lookupUsers()
}
