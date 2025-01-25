package shipfactory

import "github.com/nelsonsaake/learn-abstract-factory-design-pattern/src/ship"

type ShipFactory interface {
	NewShip()
	LaunchShip()
}

type shipfactory struct{}

func (sf shipfactory) LaunchShip(shipType string) ship.EnemyShip {
	ship := sf.NewShip(shipType)
	ship.Assemble()
	ship.DisplayEnemyShip()
	ship.FollowHeroShip()
	ship.EnemyShipShoots()
	return ship
}
