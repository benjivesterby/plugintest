package main

import (
	"context"
	"fmt"

	"github.com/benji-vesterby/atomizer"
)

// MYATOM is the exported variable for testing the Lookup method from the plugin library
var MYATOM MyAtom

func main() {
	MYATOM = MyAtom{}
}

// MyAtom is the atom struct used for testing
type MyAtom struct {
	payload []byte
}

// ID test method
func (ma *MyAtom) ID() string {
	return "testatom"
}

// Process test method
func (ma *MyAtom) Process(ctx context.Context, electron atomizer.Electron, outbound chan<- atomizer.Electron) (<-chan []byte, <-chan error) {
	var results = make(chan []byte)
	var errors = make(chan error)

	go func() {

		fmt.Println("HELLO WORLD")

	}()

	return results, errors
}
