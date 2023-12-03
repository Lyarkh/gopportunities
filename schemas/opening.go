package schemas

import (
	"gorm.io/gorm"
)

type Opening struct {
	gorm.Model
	Role string
	
}
