package main

import "fmt"

type USB interface {
	Connect()
	Name() string
}

type PhoneConnecter struct {
	name string
}

func (pc PhoneConnecter) Name() string {
	return pc.name
}
func (pc PhoneConnecter) Connect() {
	fmt.Println("device", pc.name, "is already connect")
}

func main() {
	var usb USB = PhoneConnecter{name: "ipad"}
	usb.Connect()
	Disconnect(usb)
	var a int = 0
	Disconnect(a)
}

func Disconnect(usb interface{}) {
	switch v := usb.(type) {
	case PhoneConnecter:
		fmt.Println("Disconnected:", v.name)
	case int:
		fmt.Println("this is a int")
	default:
		fmt.Println("Unkowm device")
	}
}
