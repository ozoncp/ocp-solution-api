// Package main is an entry point for module
package main

import (
	"fmt"
)

// Function main is an entry point for executable
func main() {
	const (
		description string = `
Welcome to Solution microservice!
Author: Aleksandr Fedorov
Start date: 2021/05/13
`
	)
	fmt.Print(description)
}
