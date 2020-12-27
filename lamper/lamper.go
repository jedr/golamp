package main

import (
	"fmt"
	"log"

	"jedr.github.io/lamp"
)

func main() {
	log.SetPrefix("lamper: ")

	output, err := lamp.LightUp("")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(output)
}
