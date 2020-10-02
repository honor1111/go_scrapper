package main

import (
	"fmt"

	"github.com/honor1111/mydict"
)

func main() {
	dictionary := mydict.Dictionary{}
	baseWord := "hello"
	dictionary.Add(baseWord, "First")
	err := dictionary.Update(baseWord, "Second")
	if err != nil {
		fmt.Println(err)
	}
	word, _ := dictionary.Search(baseWord)
	fmt.Println(word)
	dictionary.Delete(word)

	word2, err3 := dictionary.Search(word)
	if err3 != nil {
		fmt.Println(err3)
	} else {
		fmt.Println(word2)
	}

}
