package main

import "fmt"

// DoorBell represents a smart home doorbell system

type Doorbell struct {
	chime   string
	volume  int
	lightOn bool
}

// NewDoorBell creates a new DoorBell with default settings
func NewDoorBell() *Doorbell {
	return &Doorbell{
		chime:   "traditional",
		volume:  6,
		lightOn: false,
	}
}

// Ring to mock the ringing of doorBell
func (d *Doorbell) Ring() {
	fmt.Printf("The Doorbell is ringing with %s chime at volume %d!\n", d.chime, d.volume)
	if d.lightOn {
		fmt.Println("The lights are turned on")
	} else {
		fmt.Println("The lights are turned off")
	}
}

// SetChime sets the doorbell's chime sound
func (d *Doorbell) SetChime(chime string) {
	d.chime = chime
}

// AdjustVolume adjusts the doorBell's volume
func (d *Doorbell) AdjustVolume(volume int) {
	d.volume = volume
}

// IntegrateLights integrates the doorbell with porch lights
func (d *Doorbell) IntegrateLights(on bool) {
	d.lightOn = on
}

// Main Function
func main() {
	db := NewDoorBell()
	db.Ring() //doorbell with default value from NewDoorBell()
	db.SetChime("nepaliDhun")
	db.AdjustVolume(7)
	db.IntegrateLights(true)
	db.Ring() // neplaiDhun chime at volume lights on
}
