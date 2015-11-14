package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	fp, err := os.Open("test")
	if err != nil {
		log.Print(err)
	}
	defer fp.Close()

	if err := fp.Chmod(0664); err != nil {
		log.Print(err)
	}

	fmt.Printf("%T\n", fp)

	for pos, char := range "日本\x80語" {
		fmt.Printf("character %#U starts at byte position %d\n", char, pos)
	}
}
