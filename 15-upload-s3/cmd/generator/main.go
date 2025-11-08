package main

import (
	"fmt"
	"os"
)

var filesDir = "./tmp"

func main() {
	i := 0
	for {
		completeFileName := fmt.Sprintf("%s/file%d.txt", filesDir, i)
		f, err := os.Create(completeFileName)
		if err != nil {
			panic(err)
		}
		defer f.Close()
		contentFile := fmt.Sprintf("File number %d", i)
		f.WriteString(contentFile)
		i++
		if i >= 4 {
			break
		}
	}
}
