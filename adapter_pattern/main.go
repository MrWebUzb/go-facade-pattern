package main

import (
	"fmt"
	"strings"

	"github.com/MrWebUzb/facade/adapter_pattern/adapter"
	"github.com/MrWebUzb/facade/adapter_pattern/typec"
	"github.com/MrWebUzb/facade/adapter_pattern/usb"
)

func connect(u *usb.Usb) {
	fmt.Println("connecting mobile phone via usb to computer")
	fmt.Println("connected")
}

func main() {
	u := &usb.Usb{}
	connect(u)

	fmt.Println(strings.Repeat("-", 20))

	tpc := &typec.TypeC{}
	a := &adapter.Adapter{}

	connect(a.TypeCToUSB(tpc))
}
