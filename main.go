package main

import (
	"fmt"

	"github.com/bazeeko/onelab_homework_2/atoi"
	"github.com/bazeeko/onelab_homework_2/fibonacci"
	"github.com/bazeeko/onelab_homework_2/imports"
	"github.com/bazeeko/onelab_homework_2/itoa"
	"github.com/bazeeko/onelab_homework_2/reverse"
	"github.com/bazeeko/onelab_homework_2/runebyindex"
)

func main() {
	// reverse
	fmt.Println(reverse.Reverse("aloАЛО"))

	// itoa

	fmt.Println(itoa.Itoa(195444))

	// atoi

	num, err := atoi.Atoi("-9223372036854775808")
	if err != nil {
		fmt.Printf("%s\n", err)
	} else {
		fmt.Printf("%v\n", num)
	}

	// fibonacci

	fib := fibonacci.Fibonacci()

	for i := 0; i < 10; i++ {
		fmt.Print(fib(), " ")
	}
	fmt.Println()

	// sort imports

	err = imports.SortImports("testfile/testfile.go")
	if err != nil {
		fmt.Printf("%s\n", err)
	}

	// rune by index

	str := "aloАЛО"
	index := 6

	r, err := runebyindex.RuneByIndex(&str, &index)

	if err != nil {
		fmt.Printf("%s\n", err)
	} else {
		fmt.Printf("%c\n", r)
	}
}
