package main

import (
	"fmt"

	"github.com/honor1111/mydict"
)

func main() {
	dictionary := mydict.Dictionary{"first": "nico"}

	word := "hello"
	definition := "Greeting"

	err := dictionary.Add(word, definition)
	if err != nil {
		fmt.Println(err)
	}

	hello, _ := dictionary.Search(word)
	fmt.Println("found", word, "defintion", hello)

	err2 := dictionary.Add(word, definition)
	if err2 != nil {
		fmt.Println(err2)
	}
}
