package main

import (
	"fmt"
	"log"
	"os"
)

type Writer interface {
	Write([]byte) (int, error)
}

type File struct {
	name string
}

func (f File) Write(b []byte) (int, error) {
	file, err := os.Create(f.name)
	if err != nil {
		return 0, err
	}
	defer file.Close()
	return file.Write(b)
}

type Console struct{}

func (c Console) Write(b []byte) (int, error) {

	return fmt.Printf("%s\n", string(b))
}

func main() {
	file := File{
		name: "example.txt",
	}
	exampleText := []byte("Example Text")
	console := Console{}

	_, err := Writer.Write(file, exampleText)
	if err != nil {
		log.Fatal(err)
	}
	_, err = Writer.Write(console, exampleText)
	if err != nil {
		log.Fatal(err)
	}
}
