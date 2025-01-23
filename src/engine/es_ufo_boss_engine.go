package engine

type ESUFOBossEngine struct{}

func (e ESUFOBossEngine) GetSpeed() string {
	return "2000 mph"
}

func NewESUFOBossEngine() ESEngine {
	return &ESUFOBossEngine{}
}
