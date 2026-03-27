package projectpatterns

import "fmt"

type LightningCharger interface {
	ChargeLightning() string
}

type IPhone struct{}

func (ip *IPhone) ChargeLightning() string {
	return "Charging IPhone with lightning"
}

type USBCCharger interface {
	ChargeUSBC() string
}

type USBC struct{}

func (usbc *USBC) ChargeUSBC() string {
	return "Charging USBC with usbc"
}

type ChargeUSBCtoLightning struct {
	usbcCharge USBC
}

func (ch *ChargeUSBCtoLightning) ChargeLightning() string {
	fmt.Println("Merging USBC to Lightning")
	return ch.usbcCharge.ChargeUSBC()
}
