package main

import (
	"fmt"
	"nomard_golang/account_dict/mydict"
)

func main() {
	dictionary := mydict.Dictionary{}
	baseword := "hello"
	dictionary.Add(baseword, "First")

	dictionary.Search(baseword)
	dictionary.Delete(baseword)
	word, err := dictionary.Search(baseword)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(word)
	}
}
