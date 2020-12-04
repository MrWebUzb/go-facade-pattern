package main

type akfaBuilder struct {
	windowType string
	doorType   string
	floor      int
}

func newIglooBuilder() *akfaBuilder {
	return &akfaBuilder{}
}

func (b *akfaBuilder) setWindowType() {
	b.windowType = "Akfa Window"
}

func (b *akfaBuilder) setDoorType() {
	b.doorType = "Akfa Door"
}

func (b *akfaBuilder) setNumFloor() {
	b.floor = 3
}

func (b *akfaBuilder) getHouse() house {
	return house{
		doorType:   b.doorType,
		windowType: b.windowType,
		floor:      b.floor,
	}
}
