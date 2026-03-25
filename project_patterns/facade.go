package projectpatterns

import "fmt"

type Alarm struct{}

func (a *Alarm) SetAlarm(s string) {
	fmt.Printf("Alarm set on %s\n", s)
}

type Temp struct{}

func (t *Temp) SetTemp(s string) {
	fmt.Printf("Temp set on %s\n", s)
}

type Lights struct{}

func (l *Lights) Off() {
	fmt.Println("Lights: off")
}

type SmartHouseFacade struct {
	alarm  *Alarm
	temp   *Temp
	lights *Lights
}

func NewSmartHouseFacade() *SmartHouseFacade {
	return &SmartHouseFacade{
		alarm:  &Alarm{},
		temp:   &Temp{},
		lights: &Lights{},
	}
}

func (shf *SmartHouseFacade) GoodNight() {
	fmt.Println("Good night...")
	shf.alarm.SetAlarm("6:00")
	shf.temp.SetTemp("23")
	shf.lights.Off()
	fmt.Println("Good night activated...")
}
