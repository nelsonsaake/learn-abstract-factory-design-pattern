package weapon

type ESUFOGun struct{}

func (u ESUFOGun) GetDamage() string {
	return "20 damage"
}

func NewESUFOGun() ESWeapon {
	return &ESUFOGun{}
}
