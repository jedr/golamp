package main

import (
	"fmt"
	"time"
)

func main() {
	output := fmt.Sprintf("Current time: '%v'", time.Now().UTC())
	fmt.Println(output)
}
