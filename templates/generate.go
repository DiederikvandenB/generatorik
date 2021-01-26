package templates

import (
	"bufio"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"text/template"
)

type ComponentProps struct {
	ComponentName string
	HasProps      bool
	HasStyles     bool
}

type TestProps struct {
	ComponentName string
}

type StoryProps struct {
	ComponentName string
}

type StylesProps struct {
	ComponentName string
}

type IndexProps struct {
	ComponentName string
}

func GetTemplate(forFile string) string {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		panic(err)
	}

	basePath := ".config/generatorik/templates/"
	tplPath := filepath.Join(homeDir, basePath, forFile+".tpl")

	tpl, err := ioutil.ReadFile(tplPath)
	if err != nil {
		panic(err)
	}

	return string(tpl)
}

func WriteComponent(props ComponentProps, fileName string) {
	tpl := GetTemplate("component")

	file, close := WriteToFile(fileName)
	defer close()

	t, _ := template.New("component").Parse(tpl)
	if err := t.Execute(file, props); err != nil {
		panic(err)
	}

	if err := file.Flush(); err != nil {
		panic(err)
	}

	log.Println("Wrote component to", fileName)
}

func WriteTest(props TestProps, fileName string) {
	tpl := GetTemplate("test")

	file, close := WriteToFile(fileName)
	defer close()

	t, _ := template.New("test").Parse(tpl)
	if err := t.Execute(file, props); err != nil {
		panic(err)
	}

	if err := file.Flush(); err != nil {
		panic(err)
	}

	log.Println("Wrote test to", fileName)
}

func WriteStory(props StoryProps, fileName string) {
	tpl := GetTemplate("story")

	file, close := WriteToFile(fileName)
	defer close()

	t, _ := template.New("story").Parse(tpl)
	if err := t.Execute(file, props); err != nil {
		panic(err)
	}

	if err := file.Flush(); err != nil {
		panic(err)
	}

	log.Println("Wrote story to", fileName)
}

func WriteStyles(props StylesProps, fileName string) {
	tpl := GetTemplate("styles")

	file, close := WriteToFile(fileName)
	defer close()

	t, _ := template.New("styles").Parse(tpl)
	if err := t.Execute(file, props); err != nil {
		panic(err)
	}

	if err := file.Flush(); err != nil {
		panic(err)
	}

	log.Println("Wrote styles to", fileName)
}


func WriteIndex(props IndexProps, fileName string) {
	tpl := GetTemplate("index")

	file, close := WriteToFile(fileName)
	defer close()

	t, _ := template.New("index").Parse(tpl)
	if err := t.Execute(file, props); err != nil {
		panic(err)
	}

	if err := file.Flush(); err != nil {
		panic(err)
	}

	log.Println("Wrote index to", fileName)
}


func WriteToFile(fileName string) (*bufio.Writer, func() error) {
	file, err := os.Create(fileName)
	if err != nil {
		panic(err)
	}

	buffer := bufio.NewWriter(file)
	return buffer, file.Close
}
