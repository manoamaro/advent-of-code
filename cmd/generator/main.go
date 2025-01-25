package main

import (
	_ "embed"
	"flag"
	"fmt"
	"os"
	"text/template"

	"manoamaro.github.com/advent-of-code/pkg/errors"
)

//go:embed day.tmpl
var fs string

type Data struct {
	Year int
	Day  int
}

func main() {
	year := flag.Int("y", 0, "Year")
	day := flag.Int("d", 0, "Day")
	flag.Parse()

	if *year == 0 {
		flag.Usage()
		return
	}

	if *day == 0 {
		for i := 1; i <= 25; i++ {
			generate(*year, i)
		}
	} else {
		generate(*year, *day)
	}
}

func generate(year, day int) {
	t := template.Must(template.New("day").Parse(fs))

	data := Data{
		Year: year,
		Day:  day,
	}

	folderPath := fmt.Sprintf("cmd/%d/%d", data.Year, data.Day)
	_, err := os.Stat(folderPath)
	if !os.IsNotExist(err) {
		fmt.Println("Folder already exists", folderPath)
		return
	}
	os.MkdirAll(folderPath, 0755)
	mainFilePath := fmt.Sprintf("%s/main.go", folderPath)
	mainFile := errors.Must(os.OpenFile(mainFilePath, os.O_CREATE|os.O_WRONLY, 0644))
	defer mainFile.Close()

	fmt.Println("Generating", mainFilePath)

	err = t.Execute(mainFile, data)
	if err != nil {
		panic(err)
	}
}
