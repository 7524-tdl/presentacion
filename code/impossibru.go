package main

import (
	"fmt"
)

type musician struct {
	firstName  string
	lastName   string
	instrument string
}

// START OMIT
func main() {
	jimmy := musician{
		firstName:  "Jimmy",
		lastName:   "Page",
		instrument: "Flute",
	}
	jimmy.updateInstrument("Guitar") // HL
	fmt.Println(jimmy.firstName, "plays the", jimmy.instrument)
}

func (m musician) updateInstrument(newInstrument string) {
	m.instrument = newInstrument
}

// END OMIT
