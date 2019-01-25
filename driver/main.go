package main

import (
	"fmt"

	"github.com/slibbe/leftpad"
)

func main() {
	fmt.Println(leftpad.Format("hello", 15))
	fmt.Println(leftpad.Format("goodbye", 15))
	fmt.Println(leftpad.Format("Como está?", 15))
	fmt.Println(leftpad.Format("Internationalization", 15))
	fmt.Println(leftpad.FormatRune("hello", 15, '😊'))
	fmt.Println(leftpad.FormatRune("goodbye", 15, '😂'))
	fmt.Println(leftpad.FormatRune("Como está?", 15, '😉'))
}
