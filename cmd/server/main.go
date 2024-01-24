package main

import (
	"fmt"
	"log"
)

func Run() error {
	return nil
}

func main() {
	fmt.Println("Hello")
	if err := Run(); err != nil {
		log.Fatalln(err)
	}
}
