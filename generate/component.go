package generate

import (
  "log"
  "text/template"
)

type ComponentProps struct {
  ComponentName string
  HasProps      bool
  HasStyles     bool
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
