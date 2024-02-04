package main

import (
	"fmt"
	"os"
)

func helloWorld() {
	fmt.Println("Hello Go")
	fmt.Println("Current user:", os.Getenv("USER"))
}
