package ship

import (
	"fmt"

	"github.com/nelsonsaake/learn-abstract-factory-design-pattern/src/engine"
	"github.com/nelsonsaake/learn-abstract-factory-design-pattern/src/scf"
	"github.com/nelsonsaake/learn-abstract-factory-design-pattern/src/weapon"
)

type enemyShip struct {
	name   string
	scf    scf.ShipComponentFactory
	weapon weapon.ESWeapon
	engine engine.ESEngine
}

func (es *enemyShip) GetName() string {
	return es.name
}

func (es *enemyShip) SetName(name string) {
	es.name = name
}

func (es *enemyShip) GetShipComponentFactory() scf.ShipComponentFactory {
	return es.scf
}

func (es *enemyShip) SetShipComponentFactory(scf scf.ShipComponentFactory) {
	es.scf = scf
}

func (es *enemyShip) String() string {
	return "The " + es.GetName() + " has a top speed of " + es.engine.GetSpeed() +
		" and an attack power of " + es.weapon.GetDamage()
}

func (es *enemyShip) Setup() {
	fmt.Printf("Setting up %s\n", es.GetName())
	es.weapon = es.scf.NewWeapon()
	es.engine = es.scf.NewEngine()
}
