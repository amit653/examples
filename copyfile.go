package main

import (
	"fmt"
	"io"
	"log"
	"os"
)

func main() {

	dest := "test.txt"
	src := "defr.go"
	filein, err := os.Open(src)
	if err != nil {
		fmt.Println("D")
		log.Fatal(err)
	}
	defer filein.Close()

	fileout, err := os.Create(dest)
	if err != nil {

		log.Fatal(err)
	}
	defer fileout.Close()
	_, err = io.Copy(fileout, filein)
	if err != nil {

		log.Fatal(err)
	}
}
