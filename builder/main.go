package main

import "fmt"

func main() {
	normalBuilder := getBuilder("normal")
	akfaBuilder := getBuilder("akfa")

	director := newDirector(normalBuilder)
	normalHouse := director.buildHouse()

	fmt.Printf("Normal House Door Type: %s\n", normalHouse.doorType)
	fmt.Printf("Normal House Window Type: %s\n", normalHouse.windowType)
	fmt.Printf("Normal House Num Floor: %d\n", normalHouse.floor)

	director.setBuilder(akfaBuilder)
	akfaHouse := director.buildHouse()

	fmt.Printf("\nAkfa House Door Type: %s\n", akfaHouse.doorType)
	fmt.Printf("Akfa House Window Type: %s\n", akfaHouse.windowType)
	fmt.Printf("Akfa House Num Floor: %d\n", akfaHouse.floor)
}
