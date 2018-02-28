package basic

import (
	"fmt"
	"io/ioutil"
)

func RunFileTest() {
	files, _ := ioutil.ReadDir("/Users/fidelfly/workshop")

	for _, file := range files {
		if file.IsDir() {
			fmt.Println(file.Name())
		}
	}
}
