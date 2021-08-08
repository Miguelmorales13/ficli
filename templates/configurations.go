package templates

func ConfigDefault() string {
	return `
package configurations


func Config"%EXPORTNAME%"()  {

}
`
}
func ConfigEnv() string {
	return `
package configurations

import (
	"github.com/joho/godotenv"
	"log"
)

func ConfigEnvironments() {
	// load .env file
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

}
`
}
