package templates

func ControllersDefault() string {
	return `
package %MODELNAMEPLURAL%

type Controller struct {
}

func NewController(app fiber.Router) (c *Controller) {
	c = &Controller{}
	return

}
`
}
func ControllersCrud() string {
	return `package %MODELNAMEPLURAL%

import (
	"%PROJECTNAME%/api/%MODELNAMEPLURAL%/entities"
	"%PROJECTNAME%/models"
	"%PROJECTNAME%/utils"
	"fmt"
	"github.com/gofiber/fiber/v2"
)

type Controller struct {
	s *Service
}

func NewController(app fiber.Router) (c *Controller) {
	routes := app.Group("/%MODELNAMEPLURAL%")
	c = &Controller{NewService()}
	routes.Get("/", c.getAll)
	routes.Get("/:id", c.getOne)
	routes.Post("/", c.create)
	routes.Patch("/", c.update)
	routes.Delete("/:id", c.delete)
	return

}

// Creator%EXPORTNAME% godoc
// @Summary Create %MODELNAME%
// @Tags %MODELNAMEPLURAL%
// @Description create %MODELNAME%
// @Accept  json
// @Produce  json
// @Param user body dto.%EXPORTNAME%CreateDto true "%EXPORTNAME% create model "
// @Success 200 {object} models.ResponseSuccess{data=dto.%EXPORTNAME%UpdateDto,message=string}
// @Header 200 {string} Authorization "Bearer qwerty"
// @Failure 400,404,500 {object} models.ResponseError
// @Failure default {object} models.ResponseError
// @Router /%MODELNAMEPLURAL% [post]
func (c *Controller) create(ctx *fiber.Ctx) error {
	b := new(entities.%EXPORTNAME%)

	if err := utils.ParseBodyAndValidate(ctx, b); err != nil {
		fmt.Println(err.Error())

		return err
	}
	return ctx.JSON(nil)
}

// Update%EXPORTNAME% godoc
// @Summary Update %MODELNAME%
// @Description update %MODELNAME%
// @Accept  json
// @Produce  json
// @Tags %MODELNAMEPLURAL%
// @Param id path int true "Id from %MODELNAME% for update"
// @Param %MODELNAME% body dto.%EXPORTNAME%UpdateDto true "%EXPORTNAME% create model "
// @Success 200 {object} models.ResponseSuccess{data=dto.%EXPORTNAME%UpdateDto,message=string}
// @Header 200 {string} Authorization "Bearer qwerty"
// @Failure 400,404,500 {object} models.ResponseError
// @Failure default {object} models.ResponseError
// @Router /%MODELNAMEPLURAL%/{id} [patch]
func (c *Controller) update(ctx *fiber.Ctx) error {
	id, err := ctx.ParamsInt("id")
	if err != nil {
		return err
	}
	%MODELNAME%ToUpdate := new(entities.%EXPORTNAME%)
	if err := utils.ParseBodyAndValidate(ctx, %MODELNAME%ToUpdate); err != nil {
		fmt.Println(err.Error())
		return err
	}
	update, err := c.s.update(*%MODELNAME%ToUpdate, id)
	if err != nil {
		return err
	}
	return ctx.JSON(models.ResponseSuccess{Data: update})
}

// Delete%EXPORTNAME% godoc
// @Summary Delete %MODELNAME%
// @Description delete %MODELNAME%
// @Accept  json
// @Produce  json
// @Tags %MODELNAMEPLURAL%
// @Param id path int true "Id from %MODELNAME% for delete"
// @Success 200 {object} models.ResponseSuccess{data=int,message=string}
// @Header 200 {string} Authorization "Bearer qwerty"
// @Failure 400,404,500 {object} models.ResponseError
// @Failure default {object} models.ResponseError
// @Router /%MODELNAMEPLURAL%/{id} [delete]
func (c *Controller) delete(ctx *fiber.Ctx) error {
	id, err := ctx.ParamsInt("id")
	if err != nil {
		return err
	}
	rowsDeleted := c.s.delete(id)
	return ctx.JSON(models.ResponseSuccess{Data: rowsDeleted})

}

// List%EXPORTNAMEPLURAL% godoc
// @Summary List %MODELNAME%
// @Description get %MODELNAME%
// @Accept  json
// @Produce  json
// @Tags %MODELNAMEPLURAL%
// @Success 200 {object} models.ResponseSuccess{data=[]dto.%EXPORTNAME%UpdateDto,message=string}
// @Header 200 {string} Authorization "Bearer qwerty"
// @Failure 400,404,500 {object} models.ResponseError
// @Failure default {object} models.ResponseError
// @Router /%MODELNAMEPLURAL% [get]
func (c *Controller) getAll(ctx *fiber.Ctx) (err error) {
	%MODELNAME% := c.s.getAll()
	err = ctx.JSON(models.ResponseSuccess{Data: %MODELNAME%})
	return

}

// Get%EXPORTNAME% godoc
// @Summary Get %MODELNAME%
// @Description get %MODELNAME%
// @Accept  json
// @Produce  json
// @Tags %MODELNAMEPLURAL%
// @Param id path int true "Id from %MODELNAME% for get one"
// @Success 200 {object} models.ResponseSuccess{data=dto.%EXPORTNAME%UpdateDto,message=string}
// @Header 200 {string} Authorization "Bearer qwerty"
// @Failure 400,404,500 {object} models.ResponseError
// @Failure default {object} models.ResponseError
// @Router /%MODELNAMEPLURAL%/{id} [get]
func (c *Controller) getOne(ctx *fiber.Ctx) (err error) {
	id, err := ctx.ParamsInt("id")
	if err != nil {
		return err
	}
	%MODELNAME% := c.s.getOne(id)
	err = ctx.JSON(models.ResponseSuccess{Data: %MODELNAME%})
	return

}

`
}
