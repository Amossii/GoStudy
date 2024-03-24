package main

import "fmt"

//传递类型为值传递，为原struct的拷贝
type Usb interface {
	Connect()
}

type PhoneConnecter struct {
	Name string
}

func (pc PhoneConnecter) Connect() {
	fmt.Println("the device", pc.Name, "is connected")
}

func main() {
	var pc PhoneConnecter = PhoneConnecter{"iphone"}
	var a Usb = Usb(pc)
	a.Connect()
	pc.Name = "ipad"
	a.Connect()

	var b interface{}
	fmt.Println(b == nil)

	var p *int = nil
	b = p
	fmt.Println(b == nil)
	fmt.Println(b)
}
