package generate

import (
  "log"
  "text/template"
)

type StylesProps struct {
  ComponentName string
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
