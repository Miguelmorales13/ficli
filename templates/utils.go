package templates

func UtilsValidations() string {
	return `
package validations

import (
	"fmt"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"strings"
)

var validate = validator.New()

func Validate(payload interface{}) *fiber.Error {
	err := validate.Struct(payload)

	if err != nil {
		var errors []string
		for _, err := range err.(validator.ValidationErrors) {
			errors = append(
				errors,
				fmt.Sprintf("'%v' with value '%v' doesn't satisfy the '%v' constraint", err.Field(), err.Value(), err.Tag()),
			)
		}

		return &fiber.Error{
			Code:    fiber.StatusBadRequest,
			Message: strings.Join(errors, ","),
		}
	}

	return nil
}

// CUSTOM VALIDATION RULES =============================================

// Password validation rule: required,min=6,max=100
var _ = validate.RegisterValidation("password", func(fl validator.FieldLevel) bool {
	l := len(fl.Field().String())

	return l >= 6 && l < 100
})

`
}
func UtilsParseBody() string {
	return `
package parse_body

import (
	"%PROJECTNAME%/utils/validations"
	"github.com/gofiber/fiber/v2"
)

func ParseBody(ctx *fiber.Ctx, body interface{}) *fiber.Error {
	if err := ctx.BodyParser(body); err != nil {
		return fiber.ErrBadRequest
	}

	return nil
}

func ParseBodyAndValidate(ctx *fiber.Ctx, body interface{}) *fiber.Error {
	if err := ParseBody(ctx, body); err != nil {
		return err
	}

	return Validate(body)
}

`
}

func UtilsJWT() string {
	return `
package jwt

import (
	"errors"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"os"
	"time"
)

type TokenPayload struct {
	ID uint
}

// Generate generates the jwt token based on payload
func Generate(payload *TokenPayload) string {
	v, err := time.ParseDuration(os.Getenv("TOKEN_EXPIRATION"))
	if err != nil {
		panic("Invalid time duration. Should be time.ParseDuration string")
	}
	t := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"exp": time.Now().Add(v).Unix(),
		"ID":  payload.ID,
	})
	token, err := t.SignedString([]byte(os.Getenv("TOKEN_SECRET")))
	if err != nil {
		panic(err)
	}
	return token
}

func parse(token string) (*jwt.Token, error) {
	return jwt.Parse(token, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", t.Header["alg"])
		}
		return []byte(os.Getenv("TOKEN_SECRET")), nil
	})
}

// Verify verifies the jwt token against the secret
func Verify(token string) (*TokenPayload, error) {
	parsed, err := parse(token)
	if err != nil {
		return nil, err
	}
	claims, ok := parsed.Claims.(jwt.MapClaims)
	if !ok {
		return nil, err
	}
	id, ok := claims["ID"].(float64)
	if !ok {
		return nil, errors.New("Something went wrong")
	}
	return &TokenPayload{ID: uint(id)}, nil
}

`
}
func UtilsEmail()  string {
	return `package email

import (
	"bytes"
	template2 "html/template"
	"net/smtp"
	"os"
)

var util *EmailUtil

type EmailUtil struct {
}

func Util() *EmailUtil {
	if util == nil {
		util = &EmailUtil{}
	}
	return util
}

func (u *EmailUtil) GenTemplate(render string, options interface{}) (template string, err error) {
	var t *template2.Template
	t, err = t.ParseFiles("utils/email/templates/" + render + ".email.html")
	if err != nil {
		return
	}
	buff := new(bytes.Buffer)
	err = t.Execute(buff, options)
	if err != nil {
		return
	}
	template = buff.String()
	return
}
func (u EmailUtil) SendEmail(to []string, template string, subject string) (err error) {
	auth := smtp.PlainAuth("", os.Getenv("EMAIL_FROM"), os.Getenv("EMAIL_PASSWORD"), os.Getenv("EMAIL_HOST"))
	mime := "MIME-version: 1.0;\nContent-Type: text/html; charset=\"UTF-8\";\n\n"
	msg := "Subject: " + subject + "\n" + mime + "\n\n" + template
	err = smtp.SendMail(os.Getenv("EMAIL_HOST")+":"+os.Getenv("EMAIL_PORT"), auth, os.Getenv("EMAIL_FROM"), []string{"cacahuatisimo13@protonmail.com"}, []byte(msg))
	return
}
`
}
