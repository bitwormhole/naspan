package nasyun

import (
	nasyun "github.com/bitwormhole/nasyun"
	"github.com/bitwormhole/nasyun/gen/main4nasyun"
	"github.com/bitwormhole/nasyun/gen/test4nasyun"
	"github.com/starter-go/application"
	"github.com/starter-go/libgin/modules/libgin"
	"github.com/starter-go/libgorm/modules/libgorm"
	"github.com/starter-go/module-gorm-mysql/modules/mysql"
	"github.com/starter-go/module-gorm-sqlserver/modules/sqlserver"
	"github.com/starter-go/security/modules/security"
)

// Module  ...
func Module() application.Module {
	mb := nasyun.NewMainModule()
	mb.Components(main4nasyun.ExportComponents)

	mb.Depend(libgin.Module())
	mb.Depend(libgorm.Module())
	mb.Depend(security.Module())

	mb.Depend(mysql.Module())
	mb.Depend(sqlserver.Module())

	return mb.Create()
}

// ModuleForTest ...
func ModuleForTest() application.Module {
	mb := nasyun.NewTestModule()
	mb.Components(test4nasyun.ExportComponents)
	return mb.Create()
}
