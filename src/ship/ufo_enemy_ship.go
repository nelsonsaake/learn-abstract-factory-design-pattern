package ship

import (
	"fmt"

	"github.com/nelsonsaake/learn-abstract-factory-design-pattern/src/scf"
)

type UFOEnemyShip struct {
	enemyShip
	scf scf.ShipComponentFactory
}

func (s *UFOEnemyShip) Assemble() {
	fmt.Print("\n\n", "Assemblying Enemy Ship ", s.GetName(), "\n\n")

	s.weapon = s.scf.NewWeapon()
	s.engine = s.scf.NewEngine()
}

func NewUFOEnemyShip(shipComponentFactory scf.ShipComponentFactory) *UFOEnemyShip {
	return &UFOEnemyShip{
		scf: shipComponentFactory,
	}
}
