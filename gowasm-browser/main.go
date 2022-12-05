package main

import (
	"syscall/js"
	"time"
)

func main() {
	c := make(chan struct{})
	js.Global().Set("CurrentTime", js.FuncOf(getTime))
	<-c
}

func getTime(this js.Value, args []js.Value) interface{} {
	return time.Now().UTC().String()
}
