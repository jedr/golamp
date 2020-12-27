package main

import (
	"fmt"
	"log"
	"os"

	"jedr.github.io/lamp"
)

func main() {
	log.SetPrefix("lamper: ")

	var color string
	if len(os.Args) > 1 {
		color = os.Args[1]
	}
	log.Printf("color: '%v'", color)

	output, err := lamp.LightUp(color)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(output)
}
