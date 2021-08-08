package services

import "github.com/gertd/go-pluralize"

var libraries = []string{
	"github.com/gofiber/fiber/v2",
	"github.com/arsmn/fiber-swagger/v2",
	"github.com/go-playground/validator/v10",
	"github.com/go-openapi/swag",
	"github.com/joho/godotenv",
	"gorm.io/gorm",
	"gorm.io/driver/postgres",
}
var librariesMysql = []string{
	"gorm.io/gorm",
	"gorm.io/driver/postgres",
}
var librariesPostgres = []string{
	"gorm.io/gorm",
	"gorm.io/driver/postgres",
}
var Plural = pluralize.NewClient()

func FilesCompletes(projectName string, db string) {

	ExecuteCommand("go", "mod", "init", projectName)
	TemplateMainConfiguration(projectName)
	GetConfigProject()
	TemplateMainComplete(projectName)
	TemplateMainRunner()
	TemplateConfigurationsEnv()
	TemplateMiddlewareIndex()
	TemplateMiddlewareLogger()
	TemplateHandlerError()
	TemplateResponseError()
	TemplateResponseSuccess()
	TemplateUtilsParseBody()
	TemplateUtilsValidations()
	TemplateProvidersPostgres()
	librariesToExecute := libraries
	if db == "my" {
		librariesToExecute = append(librariesToExecute, librariesMysql...)
	} else {
		librariesToExecute = append(librariesToExecute, librariesPostgres...)
	}
	for _, library := range librariesToExecute {
		ExecuteCommand("go", "get", "-u", library)
	}
	ExecuteCommand("go", "mod", "tidy")

}
func FilesSimple(projectName string) {
	ExecuteCommand("go", "mod", "init", projectName)
	TemplateMainDefault("main")
	ExecuteCommand("go", "get", "-u", libraries[0])
}
