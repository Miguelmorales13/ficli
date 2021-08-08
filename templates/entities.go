package templates

func EntitiesDefault() string {
	return `
package entities

import (
	"gorm.io/gorm"
)

type %EXPORTNAME% struct {
	gorm.Model
}

`
}
