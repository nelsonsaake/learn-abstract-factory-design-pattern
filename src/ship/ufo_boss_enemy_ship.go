package ship

import (
	"fmt"

	"github.com/nelsonsaake/learn-abstract-factory-design-pattern/src/scf"
)

type UFOBossEnemyShip struct {
	enemyShip
	scf scf.ShipComponentFactory
}

func (s *UFOBossEnemyShip) Assemble() {
	fmt.Print("\n\n", "Assemblying Enemy Ship ", s.GetName(), "\n\n")

	s.weapon = s.scf.NewWeapon()
	s.engine = s.scf.NewEngine()
}

func NewUFOBossEnemyShip(shipComponentFactory scf.ShipComponentFactory) *UFOBossEnemyShip {
	return &UFOBossEnemyShip{
		scf: shipComponentFactory,
	}
}
