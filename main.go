package main

import (
	"nomard/mydict"
)

func main() {
	dictionary := mydict.Dictionary{"first": "First word"}
	dictionary.Add()
}
