package generate

import (
  "bufio"
  "fmt"
  "io/ioutil"
  "log"
  "os"
  "path/filepath"
  "strings"
)

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

func WriteToFile(fileName string) (*bufio.Writer, func() error) {
	file, err := os.Create(fileName)
	if err != nil {
		panic(err)
	}

	buffer := bufio.NewWriter(file)
	return buffer, file.Close
}

func OutputDir(flagValue string) string {
  var output string
  path, _ := os.Getwd()

  if flagValue != "" {
    output = flagValue
  } else {
    output = path
  }

  _, err := os.Stat(output)
  if os.IsNotExist(err) {
    reader := bufio.NewReader(os.Stdin)
    fmt.Println("Directory does not exist, do you want to create it? (y\\n)")

    create, _ := reader.ReadString('\n')
    create = strings.ToLower(create)

    if strings.Contains(create, "y") {
      err := os.MkdirAll(output, 0755)
      if err != nil {
        panic(err)
      }

      log.Println("Created directory", output)
    }
  }

  return output
}

func ComponentName(args []string, outputDir string) string {
  var componentName string
  if len(args) == 0 {
    split := strings.Split(filepath.Dir(outputDir), "/")
    componentName = split[len(split)-1]
  } else {
    componentName = args[0]
  }

  return componentName
}
