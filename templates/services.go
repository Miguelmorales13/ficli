package templates

func ServicesDefault() string {
	return `
package %MODELNAMEPLURAL%

type Service struct {
}

func NewService() *Service {
	return &Service{}
}
`
}
func ServicesCrud() string {
	return `
package %MODELNAMEPLURAL%

import (
	entities "%PROJECTNAME%/api/%MODELNAMEPLURAL%/entities"
	"%PROJECTNAME%/providers"
	"github.com/gofiber/fiber/v2"
)

type Service struct {
}

func NewService() *Service {
	return &Service{}
}
func (s *Service) create(%MODELNAME% entities.%EXPORTNAME%) (*entities.%EXPORTNAME%, error) {
	create := providers.Database.Create(&%MODELNAME%)
	return &%MODELNAME%, create.Error
}

func (s *Service) update(%MODELNAME%Updated entities.%EXPORTNAME%, id int) (*entities.%EXPORTNAME%, error) {
	%MODELNAME% := s.getOne(id)
	if %MODELNAME%.ID == 0 {
		return &entities.%EXPORTNAME%{}, fiber.NewError(fiber.StatusNotFound, "%MODELNAME% not found")
	}
	updates := providers.Database.Model(&%MODELNAME%).Updates(%MODELNAME%Updated)
	return &%MODELNAME%, updates.Error
}
func (s *Service) delete(id int) int64 {
	%MODELNAME% := s.getOne(id)
	if %MODELNAME%.ID == 0 {
		return 0
	}
	deleted := providers.Database.Delete(&%MODELNAME%)
	return deleted.RowsAffected
}
func (s *Service) getAll() (%MODELNAMEPLURAL% []entities.%EXPORTNAME%) {
	providers.Database.Find(&%MODELNAMEPLURAL%)
	return
}

func (s *Service) getOne(id int) (%MODELNAME% entities.%EXPORTNAME%) {
	providers.Database.Find(&%MODELNAME%, id)
	return
}
`
}
