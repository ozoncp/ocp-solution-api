// TODO add ocp-solution-api executable description
package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"time"
)

const (
	description string = `
Welcome to Solution microservice!
Author: Aleksandr Fedorov
Start date: 2021/05/13
`
	configFileName = `config`
)

// watchConfigFile function implements a way to watch config file updates
// TODO: add test
func watchConfigFile() {
	readConfig := func() {
		file, err := os.Open(configFileName)
		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Println(configFileName, "opened")
			// defer moves file close to the end of the function
			defer func() {
				err := file.Close()
				if err != nil {
					fmt.Println(err)
				} else {
					fmt.Println(configFileName, "closed")
				}
			}()
			bytes, err := ioutil.ReadAll(file)
			if err != nil {
				fmt.Println(err)
			} else {
				fmt.Printf("Content:\n%v\n", string(bytes))
			}
		}
	}

	for {
		readConfig()
		time.Sleep(5 * time.Second)
	}
}

// Function main is an entry point for executable
func main() {
	fmt.Print(description)
	watchConfigFile()
}
