package main

import (
	"flag"
	"fmt"
)

// This helper will streamline our error checks below.
func check(e error) {
	if e != nil {
		panic(e)
	}
}
func main() {

	folderPtr := flag.String("folder", "./", "a string")
	outputPtr := flag.String("output", "./foo.txt", "a string")
	flag.Parse()

	fmt.Println("Started to checking...")
	walkRecursive(*folderPtr, *outputPtr)

	fmt.Println("")
	fmt.Println("Started to normalize")
	normalize(*outputPtr)
}
