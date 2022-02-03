package wire

import "github.com/google/wire"

type Instance struct {
	Foo *Foo
	Bar *Bar
}

var SuperSet = wire.NewSet(NewFoo,NewBar)

func InitializeAllInstance() *Instance {
	wire.Build(SuperSet,Instance{})
	return &Instance{}
}
