package test4naspan
import (
    p2a9210c16 "github.com/bitwormhole/naspan/src/test/golang/unit"
     "github.com/starter-go/application"
)

// type p2a9210c16.DemoUnit in package:github.com/bitwormhole/naspan/src/test/golang/unit
//
// id:com-2a9210c169532405-unit-DemoUnit
// class:class-0dc072ed44b3563882bff4e657a52e62-Units
// alias:
// scope:singleton
//
type p2a9210c169_unit_DemoUnit struct {
}

func (inst* p2a9210c169_unit_DemoUnit) register(cr application.ComponentRegistry) error {
	r := cr.NewRegistration()
	r.ID = "com-2a9210c169532405-unit-DemoUnit"
	r.Classes = "class-0dc072ed44b3563882bff4e657a52e62-Units"
	r.Aliases = ""
	r.Scope = "singleton"
	r.NewFunc = inst.new
	r.InjectFunc = inst.inject
	return r.Commit()
}

func (inst* p2a9210c169_unit_DemoUnit) new() any {
    return &p2a9210c16.DemoUnit{}
}

func (inst* p2a9210c169_unit_DemoUnit) inject(injext application.InjectionExt, instance any) error {
	ie := injext
	com := instance.(*p2a9210c16.DemoUnit)
	nop(ie, com)

	


    return nil
}


