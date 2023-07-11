package main

import (
	"fmt"
	"os"
	"reloaded"
)

func main() {
	args := os.Args[1:]
	if len(args) != 2 {
		return
	}
	sample, result := args[0], args[1]
	if reloaded.CheckTxt(sample) || reloaded.CheckTxt(result) {
		fmt.Println("Error! Not a .txt format")
		return
	}
	if sample == result {
		fmt.Println("Error! Can't read and write at the same file")
		return
	}
	reloaded.WriteToFile(result, reloaded.ReadWholeFile(sample))
}
