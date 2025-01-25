package shipfactory

import (
	"strings"

	"github.com/nelsonsaake/learn-abstract-factory-design-pattern/src/scf"
	"github.com/nelsonsaake/learn-abstract-factory-design-pattern/src/ship"
)

type EnemyShipFactory struct {
	shipfactory
}

func (sf shipfactory) NewShip(shipType string) ship.EnemyShip {

	key := strings.ToUpper(shipType)

	switch key {
	case "UFO":
		shipComponentFactory := scf.NewUFOEnemyShipComponentFactory()
		theShip := ship.NewUFOEnemyShip(shipComponentFactory)
		theShip.SetName("UFO Grunt Ship")
		return theShip
	case "UFO BOSS":
		shipComponentFactory := scf.NewUFOBossEnemyShipComponentFactory()
		theShip := ship.NewUFOBossEnemyShip(shipComponentFactory)
		theShip.SetName("UFO Boss Ship")
		return theShip
	}

	return nil
}

func NewEnermyShipFactory() *EnemyShipFactory {
	return &EnemyShipFactory{}
}
