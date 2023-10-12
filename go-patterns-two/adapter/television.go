package adapter

type television interface {
	volumeUp() int
	volumeDown() int
	turnOn()
	turnOff()
}
