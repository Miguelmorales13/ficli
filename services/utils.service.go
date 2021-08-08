package services

import (
	"bytes"
	"fmt"
	"github.com/spf13/viper"
	"io/ioutil"
	"os"
	"os/exec"
	"strings"
)

func aUpper(b string) string {
	b = strings.ToLower(b)
	b = strings.Title(b)
	return b
}
func CreateMapDefault(fileName string, projectName string) map[string]string {
	MapForReplace := make(map[string]string)
	export := aUpper(fileName)
	model := strings.ToLower(fileName)
	if isSingular := Plural.IsSingular(export); isSingular {
		MapForReplace["%EXPORTNAME%"] = export
	} else {
		MapForReplace["%EXPORTNAME%"] = Plural.Singular(export)

	}
	if isPlural := Plural.IsPlural(export); isPlural {
		MapForReplace["%EXPORTNAMEPLURAL%"] = export
	} else {
		MapForReplace["%EXPORTNAMEPLURAL%"] = Plural.Plural(export)
	}
	if isSingular := Plural.IsSingular(model); isSingular {
		MapForReplace["%MODELNAME%"] = model
	} else {
		MapForReplace["%MODELNAME%"] = Plural.Singular(model)

	}
	if isPlural := Plural.IsPlural(model); isPlural {
		MapForReplace["%MODELNAMEPLURAL%"] = model
	} else {
		MapForReplace["%MODELNAMEPLURAL%"] = Plural.Plural(model)
	}

	MapForReplace["%SEMICOLON%"] = "`"
	MapForReplace["%PROJECTNAME%"] = projectName
	return MapForReplace
}

func CreateDefaultFile(template string, destFolderName string, nameStruct string, ext string) {
	fmt.Println("File", nameStruct, ext)
	defer fmt.Println("End file: ", nameStruct, ext)
	mapForReplace := CreateMapDefault(nameStruct, viper.GetString("project.name"))
	CreateTemplate(template, mapForReplace, destFolderName, nameStruct, ext)
}
func CreateDefaultFileWithProjectName(template string, destFolderName string, nameStruct string, ext string, projectName string) {
	fmt.Println("File", nameStruct, ext)
	defer fmt.Println("End file: ", nameStruct, ext)

	mapForReplace := CreateMapDefault(nameStruct, projectName)
	CreateTemplate(template, mapForReplace, destFolderName, nameStruct, ext)
}

func CreateTemplate(template string, mapForReplace map[string]string, dest string, fileName string, ext string) {
	content := ReplaceText(template, mapForReplace)

	CreateFolder(dest)
	CreateFile(dest+"/"+fileName+"."+ext, content)
}
func CreateFolder(directory string) {
	err := os.MkdirAll(directory, 0755)
	if err != nil {
		panic(err)
	}
}
func ReplaceText(content string, MapForReplace map[string]string) string {
	input := content
	for key, value := range MapForReplace {
		input = strings.Replace(input, key, value, -1)
	}
	return input
}
func CreateFile(dest string, data string) {
	err := ioutil.WriteFile(dest, []byte(data), 0644)
	if err != nil {
		panic(err)
	}
}
func GetConfigProject() string {
	viper.AddConfigPath(".")
	viper.SetConfigName("fiber")
	viper.SetConfigType("yml")
	viper.AutomaticEnv()
	err := viper.ReadInConfig()
	if err != nil {
		fmt.Println(err.Error())
		panic(err)
	} // Need to explicitly set this to json
	return viper.GetString("project.name")
}
func ExecuteCommand(command string, args ...string) {
	cmd := exec.Command(command, args...)
	fmt.Println("Executing:", cmd.String())
	cmdOutput := &bytes.Buffer{}
	cmd.Stdout = cmdOutput
	err := cmd.Run()
	if err != nil {
		fmt.Println("Error in ", err.Error())
	}
	fmt.Println(string(cmdOutput.Bytes()))
}
