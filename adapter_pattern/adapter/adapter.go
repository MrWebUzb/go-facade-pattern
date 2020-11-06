package adapter

import (
	"fmt"

	"github.com/MrWebUzb/facade/adapter_pattern/typec"
	"github.com/MrWebUzb/facade/adapter_pattern/usb"
)

// Adapter ...
type Adapter struct{}

// USBToTypeC ...
func (a *Adapter) USBToTypeC(u *usb.Usb) *typec.TypeC {
	fmt.Println("Got USB connector")
	defer fmt.Println("Returned TypeC connector")
	return &typec.TypeC{}
}

// TypeCToUSB ...
func (a *Adapter) TypeCToUSB(tpc *typec.TypeC) *usb.Usb {
	fmt.Println("Got TypeC connector")
	defer fmt.Println("Returned TypeC connector")
	return &usb.Usb{}
}
