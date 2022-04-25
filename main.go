package main

import (
	"fmt"

	"github.com/ppark2ya/go-basics/mydict"
)

func main() {
	dictionary := mydict.Dictionary{"fruit": "apple"}
	baseWord := "hello"
	dictionary.Add(baseWord, "First")
	// err := dictionary.Update(baseWord, "Second")

	// if err != nil {
	// 	fmt.Println(err)
	// }
	dictionary.Delete(baseWord)
	word, err := dictionary.Search(baseWord)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(word)
}
