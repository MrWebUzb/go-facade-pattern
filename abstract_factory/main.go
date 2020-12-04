package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/MrWebUzb/facade/abstract_factory/factory"
	"github.com/MrWebUzb/facade/abstract_factory/factory/brands"
)

func main() {
	if adidasFactory, err := brands.GetSportsFactory("adidas"); err == nil {
		shoe := adidasFactory.MakeShoe()
		shirt := adidasFactory.MakeShirt()
		printShoeDetails(shoe)
		printShirtDetails(shirt)
	} else {
		fmt.Fprintf(os.Stderr, "unable to connect with adidas factory: %v\n", err)
		os.Exit(1)
	}

	if nikeFactory, err := brands.GetSportsFactory("nike"); err == nil {
		shoe := nikeFactory.MakeShoe()
		shirt := nikeFactory.MakeShirt()
		printShoeDetails(shoe)
		printShirtDetails(shirt)
	} else {
		fmt.Fprintf(os.Stderr, "unable to connect with nike factory: %v\n", err)
		os.Exit(1)
	}
}

func printShoeDetails(s factory.IShoe) {
	fmt.Println("Shoe details")
	fmt.Println(strings.Repeat("-", 20))
	fmt.Printf("Logo: %s\n", s.GetLogo())
	fmt.Printf("Size: %d\n", s.GetSize())
	fmt.Println(strings.Repeat("-", 20))
}

func printShirtDetails(s factory.IShirt) {
	fmt.Println("Shirt details")
	fmt.Println(strings.Repeat("-", 20))
	fmt.Printf("Logo: %s\n", s.GetLogo())
	fmt.Printf("Size: %d\n", s.GetSize())
	fmt.Println(strings.Repeat("-", 20))
}
