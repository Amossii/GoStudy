package main

import "fmt"

type empty interface {
}
type USB interface {
	Name() string
	connecter // 嵌入接口
}
type connecter interface {
	Connect()
}
type PhoneConnecter struct {
	name string
}

func (pc PhoneConnecter) Name() string {
	return pc.name
}
func (pc PhoneConnecter) Connect() {
	fmt.Println("connect:", pc.name)
}
func main() {
	var a USB
	a = PhoneConnecter{"PhoneConnecter"}
	a.Connect()
	Disconnect(a)
}
func Disconnect(usb USB) {
	//判断接口中的结构是否是PhoneConnecter类型
	if pc, ok := usb.(PhoneConnecter); ok {
		fmt.Println("disconnect:", pc.name)
		return
	}
	fmt.Println(("Unknown device."))
}
