package dao

import (
	"github.com/bitwormhole/naspan/app/data/dxo"
	"github.com/bitwormhole/naspan/app/data/entity"
	"gorm.io/gorm"
)

// ExampleDAO ...
type ExampleDAO interface {
	Find(db *gorm.DB, id dxo.ExampleID) (*entity.Example, error)
}
