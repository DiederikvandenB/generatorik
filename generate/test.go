package generate

import (
  "log"
  "text/template"
)

type TestProps struct {
  ComponentName string
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
