package experiments

import (
	"fmt"
	"io/ioutil"
	"os"
	"time"
)

// watchConfigFile function implements a way to watch config file updates
func watchConfigFile() {
	const filename string = `filename`

	readConfig := func() {
		file, err := os.Open(filename)
		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Println(filename, "opened")
			// defer moves file close to the end of the function
			defer func() {
				err := file.Close()
				if err != nil {
					fmt.Println(err)
				} else {
					fmt.Println(filename, "closed")
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
