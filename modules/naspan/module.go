package naspan

import (
	naspan "github.com/bitwormhole/naspan"
	"github.com/bitwormhole/naspan/gen/main4naspan"
	"github.com/bitwormhole/naspan/gen/test4naspan"
	"github.com/starter-go/application"
	"github.com/starter-go/libgin/modules/libgin"
	"github.com/starter-go/libgorm/modules/libgorm"
	"github.com/starter-go/module-gorm-mysql/modules/mysql"
	"github.com/starter-go/module-gorm-sqlserver/modules/sqlserver"
	"github.com/starter-go/security/modules/security"
)

// Module  ...
func Module() application.Module {
	mb := naspan.NewMainModule()
	mb.Components(main4naspan.ExportComponents)

	mb.Depend(libgin.Module())
	mb.Depend(libgorm.Module())
	mb.Depend(security.Module())

	mb.Depend(mysql.Module())
	mb.Depend(sqlserver.Module())

	return mb.Create()
}

// ModuleForTest ...
func ModuleForTest() application.Module {
	mb := naspan.NewTestModule()
	mb.Components(test4naspan.ExportComponents)
	return mb.Create()
}
