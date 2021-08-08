package templates

func ModelsResponseError() string {
	return `
package models

import "time"

type ResponseError struct {
	Code       int       %SEMICOLON%json:"code"%SEMICOLON%
	Path       string    %SEMICOLON%json:"path"%SEMICOLON%
	Method     string    %SEMICOLON%json:"method"%SEMICOLON%
	TimeStamps time.Time %SEMICOLON%json:"timestamps"%SEMICOLON%
	Message    string    %SEMICOLON%json:"message"%SEMICOLON%
}


`
}
func ModelsResponseSuccess() string {
	return `
package models

type ResponseSuccess struct {
	Data    interface{} %SEMICOLON%json:"data"%SEMICOLON%
	Message string      %SEMICOLON%json:"message"%SEMICOLON%
}

`
}

func ModelsCreateDto() string {
	return `
package dto

type %EXPORTNAME%CreateDto struct {
	
}
`
}
func ModelsDefaultDto() string {
	return `
package dto

type %EXPORTNAME%Dto struct {
	
}
`
}
func ModelsUpdateDto() string {
	return `
package dto

type %EXPORTNAME%UpdateDto struct {
	%EXPORTNAME%CreateDto
}
`
}
