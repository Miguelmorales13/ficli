package templates

func EnvironmentsDefault() string {
	return `
PORT=:8000
`
}
func EnvironmentsProd() string {
	return `
PORT=:8000
`
}
func EnvironmentsTest() string {
	return `
PORT=:8000
`
}
