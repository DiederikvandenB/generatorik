package generate

import (
  "log"
  "text/template"
)

type ContextProps struct {
  ContextName string
}

func WriteContext(props ContextProps, fileName string) {
  tpl := GetTemplate("context")

  file, close := WriteToFile(fileName)
  defer close()

  t, _ := template.New("context").Parse(tpl)
  if err := t.Execute(file, props); err != nil {
    panic(err)
  }

  if err := file.Flush(); err != nil {
    panic(err)
  }

  log.Println("Wrote context to", fileName)
}
