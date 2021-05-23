// Package experiments contains experimental functions checking Go behaviour
package experiments

import (
	"fmt"
	"os"
)

// DeferFileClose implements a way to defer file close
// TODO: add test
func deferFileClose() {
	filename := "tmp"
	for i := 0; i < 10; i++ {
		file, err := os.Create(filename)
		if err != nil {
			panic(err)
		}
		fmt.Println(filename, "created")
		// defer moves file close to the end of the function
		defer func() {
			err := file.Close()
			if err != nil {
				panic(err)
			}
			fmt.Println(filename, "closed")
		}()
	}
}
