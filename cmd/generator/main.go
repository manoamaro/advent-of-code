package main

import (
	_ "embed"
	"fmt"
	"os"
	"strconv"
	"text/template"
)

//go:embed day.tmpl
var fs string

type Data struct {
	Day int
}

func main() {

	argDay := os.Args[1]

	day, err := strconv.Atoi(argDay)
	if err != nil {
		panic(err)
	}

	data := Data{
		Day: day,
	}

	t := template.Must(template.New("day").Parse(fs))

	folderPath := fmt.Sprintf("cmd/day%d", data.Day)
	os.MkdirAll(folderPath, 0755)
	mainFilePath := fmt.Sprintf("%s/main.go", folderPath)
	mainFile, err := os.OpenFile(mainFilePath, os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		panic(err)
	}

	defer mainFile.Close()

	err = t.Execute(mainFile, data)
	if err != nil {
		panic(err)
	}
}
