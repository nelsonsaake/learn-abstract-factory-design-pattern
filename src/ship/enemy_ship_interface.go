package ship

type EnemyShip interface {
	FollowHeroShip()
	DisplayEnemyShip()
	EnemyShipShoots()
	String() string
}
