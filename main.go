package main

import (
	"fmt"
	"os"
)

func init() {
	// sets up global context & is executed before main()
	fmt.Print("Setting up Global Context")
}

func main() {
	fmt.Println("\nHello Go")
	fmt.Println("Current user:", os.Getenv("USER"))
}
