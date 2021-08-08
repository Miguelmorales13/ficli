package templates

func InitDefault() string {
	return `
package main

import (
	"github.com/gofiber/fiber/v2"
	"log"
)

func main() {
	app := fiber.New()

	log.Fatal(app.Listen(":8000"))
}

`
}
func InitRunner() string {
	return `
root:              .
tmp_path:          ./tmp
build_name:        runner-build
build_log:         runner-build-errors.log
valid_ext:         .go, .tpl, .tmpl, .html
no_rebuild_ext:    .tpl, .tmpl, .html
ignored:           assets, tmp
build_delay:       600
colors:            1
log_color_main:    cyan
log_color_build:   yellow
log_color_runner:  green
log_color_watcher: magenta
log_color_app:
`
}
func InitConfig() string {
	return `project:
  name: "%PROJECTNAME%"

`
}
func InitComplete() string {
	return `
package main

import (
	"%PROJECTNAME%/api/users"
	"%PROJECTNAME%/configurations"
	_ "%PROJECTNAME%/docs"
	"%PROJECTNAME%/handlers"
	"%PROJECTNAME%/middlewares"
	"%PROJECTNAME%/providers"
	swagger "github.com/arsmn/fiber-swagger/v2"
	"github.com/gofiber/fiber/v2"
	recover2 "github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/gofiber/fiber/v2/middleware/requestid"
	"log"
	"os"
)

// @title Curd
// @version 1.0
// @description This is a sample server celler server.
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:8080
// @BasePath /api
// @query.collection.format multi

// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization

// @x-extension-openapi {"example": "value on a json format"}
func main() {
	configurations.ConfigEnvironments()
	providers.InitDbPostgres()
	app := fiber.New(
		fiber.Config{
			ErrorHandler: handlers.ErrorHandler,
		},

	)

	middlewares.Config(app)
	app.Use(recover2.New())
	app.Use(requestid.New())
	app.Get("/docs/*", swagger.Handler) // default

	app.Get("/docs/*", swagger.New(swagger.Config{ // custom
		URL:         "http://example.com/doc.json",
		DeepLinking: false,
	}))
	_ := app.Group("/api")

	log.Fatal(app.Listen(os.Getenv("PORT")))
}

`
}
