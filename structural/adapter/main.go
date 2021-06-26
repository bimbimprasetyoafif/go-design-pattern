package main

import "fmt"

type Human struct {}

type Computer interface {
	InsertIntoPortUSB()
}

func (h *Human) InsertIntoPort(com Computer) {
	fmt.Println("human insert keyboard")
	com.InsertIntoPortUSB()
}

type ModernComputer struct {}

func (m *ModernComputer) InsertIntoPortUSB() {
	fmt.Println("usb plug active")
}

type OldComputer struct {}

func (m *OldComputer) InsertIntoPortPS2() {
	fmt.Println("ps2 plug active")
}

type USBToPS2Adapter struct {
	Computer *OldComputer
}

func (u *USBToPS2Adapter) InsertIntoPortUSB() {
	fmt.Println("using converter usb to ps2")
	u.Computer.InsertIntoPortPS2()
}

func main() {
	me := &Human{}
	modernCom := &ModernComputer{}
	oldCom := &OldComputer{}

	me.InsertIntoPort(modernCom)

	adapterPS2 := &USBToPS2Adapter{
		Computer: oldCom,
	}
	me.InsertIntoPort(adapterPS2)

}
