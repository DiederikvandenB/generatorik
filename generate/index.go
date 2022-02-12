package generate

import (
  "log"
  "text/template"
)

type IndexProps struct {
  ComponentName string
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
