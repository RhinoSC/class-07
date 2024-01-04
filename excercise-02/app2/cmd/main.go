package main

import (
	"fmt"
	"io"
	"os"
)

func main() {

	path := "./cmd/customer.txt"

	text := ReadFile(path)

	fmt.Println(text)
	fmt.Println("Ejecucion finalizada")

}

func ReadFile(path string) string {
	file, err := os.Open(path)

	defer func() {
		err := recover()
		if err != nil {
			fmt.Println(err)
		}
	}()

	if err != nil {
		panic("The indicated file was not found or is damaged")
	}

	// read bytes from file
	bytes, err := io.ReadAll(file)
	if err != nil {
		panic("An error occurred while reading the file")
	}

	defer file.Close()
	return string(bytes)
}
