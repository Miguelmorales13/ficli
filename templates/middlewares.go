package templates

func MiddlewaresDefault() string {
	return `
package middlewares

import (
	"github.com/gofiber/fiber/v2"
)

func %EXPORTNAME%Middleware(e *fiber.App) {
}

`
}
func MiddlewaresIndex() string {
	return `
package middlewares

import (
	"github.com/gofiber/fiber/v2"
)

func Config(e *fiber.App) {
}

`
}
func MiddlewaresLogger() string {
	return `
package middlewares

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"os"
)

func LoggerMiddleware() fiber.Handler {
	return logger.New(logger.Config{
		Format:     "${red}[${time}]-${blue}[${locals:requestid}] [${ip}|${latency}] [${method}|${url}] [${status}|${body}]  \n",
		TimeFormat: "2006-01-02 15:04:05.00000",
		Output:     os.Stderr,
	})
}
`
}
