package main

import (
	"fmt"

	"github.com/nelsonsaake/learn-abstract-factory-design-pattern/src/shipfactory"
)

func main() {

	sf := shipfactory.NewEnermyShipFactory()

	grunt := sf.LaunchShip("UFO")
	fmt.Println(grunt)

	boss := sf.LaunchShip("UFO BOSS")
	fmt.Println(boss)
}
