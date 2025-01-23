package engine

type ESUFOEngine struct{}

func (e ESUFOEngine) GetSpeed() string {
	return "1000 mph"
}

func NewESUFOEngine() ESEngine {
	return &ESUFOEngine{}
}
