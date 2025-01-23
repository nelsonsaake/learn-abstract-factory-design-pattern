package scf

import (
	"github.com/nelsonsaake/learn-abstract-factory-design-pattern/src/engine"
	"github.com/nelsonsaake/learn-abstract-factory-design-pattern/src/weapon"
)

type UFOBossEnemyShipComponentFactory struct{}

func (f UFOBossEnemyShipComponentFactory) NewWeapon() weapon.ESWeapon {
	return weapon.NewESUFOBossGun()
}

func (f UFOBossEnemyShipComponentFactory) NewEngine() engine.ESEngine {
	return engine.NewESUFOBossEngine()
}

func NewUFOBossEnemyShipComponentFactory() ShipComponentFactory {
	return &UFOBossEnemyShipComponentFactory{}
}
