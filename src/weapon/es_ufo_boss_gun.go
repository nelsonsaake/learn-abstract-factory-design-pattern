package weapon

type ESUFOBossGun struct{}

func (u ESUFOBossGun) GetDamage() string {
	return "40 damage"
}

func NewESUFOBossGun() ESWeapon {
	return &ESUFOBossGun{}
}
