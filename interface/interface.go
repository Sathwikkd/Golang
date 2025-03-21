package main

import "fmt"

type operateDevice interface {
	connect()
	disconnect()
}

type wifi struct {
	name     string
	password string
}

type bluetooth struct {
	name string
}

func (w wifi) connect() {
	fmt.Println(w.name, "connected to wifi through ", w.password)
}

func (w wifi) disconnect() {
	fmt.Println(w.name, "disconnected")
}

func (b bluetooth) connect() {
	fmt.Println(b.name, "connected to bluetooth ")
}
func (b bluetooth) disconnect() {
	fmt.Println(b.name, "disconnected")
}

func Device(d operateDevice) {
	d.connect()
	d.disconnect()
}

func main() {
	wifi_device := wifi{name: "rust", password: "1234"}
	bluetooth_device := bluetooth{name: "airdopes 131"}

	Device(wifi_device)
	Device(bluetooth_device)
}
