// scf: Ship Component Factory
package scf

import (
	"github.com/nelsonsaake/learn-abstract-factory-design-pattern/src/engine"
	"github.com/nelsonsaake/learn-abstract-factory-design-pattern/src/weapon"
)

type ShipComponentFactory interface {
	NewWeapon() weapon.ESWeapon
	NewEngine() engine.ESEngine
}
