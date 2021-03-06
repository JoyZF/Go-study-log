package _package

type AnimalSounder interface {
	MakeDNA()
}



type Animal struct {
	name string
}

type Cat struct {
	// 继承
	Animal
	FeatureA string
}

type Dog struct {
	Animal
	FeatureB string
}

func NewAnimal() *Animal  {
	return &Animal{}
}

func (p *Animal) SetName(name string)  {
	// 封装
	p.name = name
}

func (p *Animal) GetName() string {
	return p.name
}

func MakeSomeDna(sounder AnimalSounder)  {
	sounder.MakeDNA()
}