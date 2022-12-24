package creational

import (
	"fmt"
)

type driver struct {
	name string
}

var carDriver *driver

// NewDriver create new car driver.
func NewDriver(n string) *driver {
	if carDriver != nil {
		return nil
	}
	carDriver = &driver{n}
	fmt.Printf("%s driver Created.\n", carDriver.name)
	return carDriver
}
