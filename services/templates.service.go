package services

import (
	"github.com/Miguelmorales13/ficli/templates"
)

func TemplateControllerDefault(destFileName string) {
	CreateDefaultFile(templates.ControllersDefault(), "api/"+destFileName, destFileName, "controller.go")
}
func TemplateControllerCrud(destFileName string) {
	CreateDefaultFile(templates.ControllersCrud(), "api/"+destFileName, destFileName, "controller.go")
}

func TemplateConfigurationsEnv() {
	CreateDefaultFile(templates.ConfigEnv(), "configurations", "env", "config.go")
	TemplateEnvironments()
	ExecuteCommand("go", "get", "-u", "github.com/joho/godotenv")
}
func TemplateConfigurationDefault(destFileName string) {
	CreateDefaultFile(templates.ConfigDefault(), "configurations", destFileName, "config.go")
}
func TemplateEntity(destFileName string) {
	CreateDefaultFile(templates.EntitiesDefault(), "api/"+destFileName+"/entities", Plural.Singular(destFileName), "entity.go")
}
func TemplateEnvironments() {
	CreateDefaultFile(templates.EnvironmentsDefault(), ".", "", "env")
	CreateDefaultFile(templates.EnvironmentsProd(), ".", "", "prod.env")
	CreateDefaultFile(templates.EnvironmentsTest(), ".", "", "test.env")
}
func TemplateHandlerError() {
	CreateDefaultFile(templates.HandlersError(), "handlers", "error", "handler.go")
}
func TemplateHandlerDefault(destFileName string) {
	CreateDefaultFile(templates.HandlersDefault(), "handlers", destFileName, "handler.go")
}
func TemplateMainComplete(projectName string) {
	CreateDefaultFileWithProjectName(templates.InitComplete(), ".", "main", "go", projectName)
}
func TemplateMainConfiguration(projectName string) {
	CreateDefaultFileWithProjectName(templates.InitConfig(), ".", "fiber", "yml", projectName)
}
func TemplateMainDefault(projectName string) {
	CreateDefaultFileWithProjectName(templates.InitDefault(), ".", "main", "go", projectName)
}
func TemplateMainRunner() {
	CreateDefaultFile(templates.InitRunner(), ".", "runner", "conf")
}
func TemplateMiddlewareIndex() {
	CreateDefaultFile(templates.MiddlewaresIndex(), "middlewares", "index", "go")
}
func TemplateMiddlewareDefault(destFileName string) {
	CreateDefaultFile(templates.MiddlewaresDefault(), "middlewares", destFileName, "middleware.go")
}
func TemplateMiddlewareLogger() {
	CreateDefaultFile(templates.MiddlewaresLogger(), "middlewares", "logger", "middleware.go")
}
func TemplateResponseError() {
	CreateDefaultFile(templates.ModelsResponseError(), "models", "responseError", "model.go")
}
func TemplateCreateDto(destFileName string) {
	CreateDefaultFile(templates.ModelsCreateDto(), "api/"+destFileName+"/dto", Plural.Singular(destFileName), "create.dto.go")
}
func TemplateUpdateDto(destFileName string) {
	CreateDefaultFile(templates.ModelsUpdateDto(), "api/"+destFileName+"/dto", Plural.Singular(destFileName), "update.dto.go")
}
func TemplateDefaultDto(destFileName string) {
	CreateDefaultFile(templates.ModelsDefaultDto(), "api/"+destFileName+"/dto", Plural.Singular(destFileName), "dto.go")
}
func TemplateResponseSuccess() {
	CreateDefaultFile(templates.ModelsResponseSuccess(), "models", "responseSuccess", "model.go")
}
func TemplateServiceDefault(destFileName string) {
	CreateDefaultFile(templates.ServicesDefault(), "api/"+destFileName, destFileName, "service.go")
}
func TemplateServiceCrud(destFileName string) {
	CreateDefaultFile(templates.ServicesCrud(), "api/"+destFileName, destFileName, "service.go")
}

func TemplateUtilsParseBody() {
	CreateDefaultFile(templates.UtilsParseBody(), "utils/parse_body", "parse_body", "util.go")
}

func TemplateUtilsValidations() {
	CreateDefaultFile(templates.UtilsValidations(), "utils/validations", "validations", "util.go")
}
func TemplateUtilsJWT() {
	CreateDefaultFile(templates.UtilsJWT(), "utils/jwt", "jwt", "util.go")
}
func TemplateUtilsEmails() {
	CreateDefaultFile(templates.UtilsEmail(), "utils/email", "email", "util.go")
}
func TemplateProvidersMysql() {
	CreateDefaultFile(templates.ProvidersGormMysql(), "providers", "gorm-mysql", "provider.go")
	ExecuteCommand("go","get","-u","github.com/dgrijalva/jwt-go")
}
func TemplateProvidersPostgres() {
	CreateDefaultFile(templates.ProvidersGormPostgres(), "providers", "gorm-postgres", "provider.go")
}
