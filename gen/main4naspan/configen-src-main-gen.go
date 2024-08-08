package main4naspan
import (
    p10eb65bc6 "github.com/bitwormhole/naspan/app/data/dao"
    p1e004ce56 "github.com/bitwormhole/naspan/app/data/database"
    p2f7744998 "github.com/bitwormhole/naspan/app/data/dxo"
    pb686eecd9 "github.com/bitwormhole/naspan/app/implements/example"
    p6f0cf9a44 "github.com/bitwormhole/naspan/app/web/controllers/apiv1"
    pd1a916a20 "github.com/starter-go/libgin"
    p512a30914 "github.com/starter-go/libgorm"
     "github.com/starter-go/application"
)

// type p1e004ce56.ThisGroup in package:github.com/bitwormhole/naspan/app/data/database
//
// id:com-1e004ce56683b2ff-database-ThisGroup
// class:class-512a309140d0ad99eb1c95c8dc0d02f9-GroupRegistry
// alias:alias-2f77449983018ad3c72066a59c23901c-DatabaseAgent
// scope:singleton
//
type p1e004ce566_database_ThisGroup struct {
}

func (inst* p1e004ce566_database_ThisGroup) register(cr application.ComponentRegistry) error {
	r := cr.NewRegistration()
	r.ID = "com-1e004ce56683b2ff-database-ThisGroup"
	r.Classes = "class-512a309140d0ad99eb1c95c8dc0d02f9-GroupRegistry"
	r.Aliases = "alias-2f77449983018ad3c72066a59c23901c-DatabaseAgent"
	r.Scope = "singleton"
	r.NewFunc = inst.new
	r.InjectFunc = inst.inject
	return r.Commit()
}

func (inst* p1e004ce566_database_ThisGroup) new() any {
    return &p1e004ce56.ThisGroup{}
}

func (inst* p1e004ce566_database_ThisGroup) inject(injext application.InjectionExt, instance any) error {
	ie := injext
	com := instance.(*p1e004ce56.ThisGroup)
	nop(ie, com)

	
    com.Enabled = inst.getEnabled(ie)
    com.Alias = inst.getAlias(ie)
    com.URI = inst.getURI(ie)
    com.Prefix = inst.getPrefix(ie)
    com.Source = inst.getSource(ie)
    com.SourceManager = inst.getSourceManager(ie)


    return nil
}


func (inst*p1e004ce566_database_ThisGroup) getEnabled(ie application.InjectionExt)bool{
    return ie.GetBool("${datagroup.default.enabled}")
}


func (inst*p1e004ce566_database_ThisGroup) getAlias(ie application.InjectionExt)string{
    return ie.GetString("${datagroup.default.alias}")
}


func (inst*p1e004ce566_database_ThisGroup) getURI(ie application.InjectionExt)string{
    return ie.GetString("${datagroup.default.uri}")
}


func (inst*p1e004ce566_database_ThisGroup) getPrefix(ie application.InjectionExt)string{
    return ie.GetString("${datagroup.default.table-name-prefix}")
}


func (inst*p1e004ce566_database_ThisGroup) getSource(ie application.InjectionExt)string{
    return ie.GetString("${datagroup.default.datasource}")
}


func (inst*p1e004ce566_database_ThisGroup) getSourceManager(ie application.InjectionExt)p512a30914.DataSourceManager{
    return ie.GetComponent("#alias-512a309140d0ad99eb1c95c8dc0d02f9-DataSourceManager").(p512a30914.DataSourceManager)
}



// type pb686eecd9.DaoImpl in package:github.com/bitwormhole/naspan/app/implements/example
//
// id:com-b686eecd91ca99f7-example-DaoImpl
// class:
// alias:alias-10eb65bc64df9604d6518395e2cf1ec0-ExampleDAO
// scope:singleton
//
type pb686eecd91_example_DaoImpl struct {
}

func (inst* pb686eecd91_example_DaoImpl) register(cr application.ComponentRegistry) error {
	r := cr.NewRegistration()
	r.ID = "com-b686eecd91ca99f7-example-DaoImpl"
	r.Classes = ""
	r.Aliases = "alias-10eb65bc64df9604d6518395e2cf1ec0-ExampleDAO"
	r.Scope = "singleton"
	r.NewFunc = inst.new
	r.InjectFunc = inst.inject
	return r.Commit()
}

func (inst* pb686eecd91_example_DaoImpl) new() any {
    return &pb686eecd9.DaoImpl{}
}

func (inst* pb686eecd91_example_DaoImpl) inject(injext application.InjectionExt, instance any) error {
	ie := injext
	com := instance.(*pb686eecd9.DaoImpl)
	nop(ie, com)

	
    com.Agent = inst.getAgent(ie)


    return nil
}


func (inst*pb686eecd91_example_DaoImpl) getAgent(ie application.InjectionExt)p2f7744998.DatabaseAgent{
    return ie.GetComponent("#alias-2f77449983018ad3c72066a59c23901c-DatabaseAgent").(p2f7744998.DatabaseAgent)
}



// type p6f0cf9a44.ExampleController in package:github.com/bitwormhole/naspan/app/web/controllers/apiv1
//
// id:com-6f0cf9a44d38deda-apiv1-ExampleController
// class:class-d1a916a203352fd5d33eabc36896b42e-Controller
// alias:
// scope:singleton
//
type p6f0cf9a44d_apiv1_ExampleController struct {
}

func (inst* p6f0cf9a44d_apiv1_ExampleController) register(cr application.ComponentRegistry) error {
	r := cr.NewRegistration()
	r.ID = "com-6f0cf9a44d38deda-apiv1-ExampleController"
	r.Classes = "class-d1a916a203352fd5d33eabc36896b42e-Controller"
	r.Aliases = ""
	r.Scope = "singleton"
	r.NewFunc = inst.new
	r.InjectFunc = inst.inject
	return r.Commit()
}

func (inst* p6f0cf9a44d_apiv1_ExampleController) new() any {
    return &p6f0cf9a44.ExampleController{}
}

func (inst* p6f0cf9a44d_apiv1_ExampleController) inject(injext application.InjectionExt, instance any) error {
	ie := injext
	com := instance.(*p6f0cf9a44.ExampleController)
	nop(ie, com)

	
    com.Sender = inst.getSender(ie)
    com.Dao = inst.getDao(ie)


    return nil
}


func (inst*p6f0cf9a44d_apiv1_ExampleController) getSender(ie application.InjectionExt)pd1a916a20.Responder{
    return ie.GetComponent("#alias-d1a916a203352fd5d33eabc36896b42e-Responder").(pd1a916a20.Responder)
}


func (inst*p6f0cf9a44d_apiv1_ExampleController) getDao(ie application.InjectionExt)p10eb65bc6.ExampleDAO{
    return ie.GetComponent("#alias-10eb65bc64df9604d6518395e2cf1ec0-ExampleDAO").(p10eb65bc6.ExampleDAO)
}


