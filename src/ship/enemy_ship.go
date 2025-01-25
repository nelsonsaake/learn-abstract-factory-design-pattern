package ship

import (
	"github.com/nelsonsaake/learn-abstract-factory-design-pattern/src/engine"
	"github.com/nelsonsaake/learn-abstract-factory-design-pattern/src/weapon"
)

type EnemyShip interface {
	Assemble()
	FollowHeroShip()
	DisplayEnemyShip()
	EnemyShipShoots()
	String() string
}

type enemyShip struct {
	name   string
	weapon weapon.ESWeapon
	engine engine.ESEngine
}

func (es *enemyShip) GetName() string {
	return es.name
}

func (es *enemyShip) SetName(name string) {
	es.name = name
}

func (es *enemyShip) FollowHeroShip() {
	println(es.GetName(), "is following the hero at", es.engine.GetSpeed())
}

func (es *enemyShip) DisplayEnemyShip() {
	println(es.GetName(), "is on the screen")
}

func (es *enemyShip) EnemyShipShoots() {
	println(es.GetName(), "attacks and does", es.weapon.GetDamage())
}

func (es *enemyShip) String() string {
	return "The " + es.GetName() + " has a top speed of " + es.engine.GetSpeed() +
		" and an attack power of " + es.weapon.GetDamage()
}
