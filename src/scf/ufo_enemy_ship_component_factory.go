package scf

import (
	"github.com/nelsonsaake/learn-abstract-factory-design-pattern/src/engine"
	"github.com/nelsonsaake/learn-abstract-factory-design-pattern/src/weapon"
)

type UFOEnemyShipComponentFactory struct{}

func (f UFOEnemyShipComponentFactory) NewWeapon() weapon.ESWeapon {
	return weapon.NewESUFOGun()
}

func (f UFOEnemyShipComponentFactory) NewEngine() engine.ESEngine {
	return engine.NewESUFOEngine()
}

func NewUFOEnemyShipComponentFactory() ShipComponentFactory {
	return &UFOEnemyShipComponentFactory{}
}
