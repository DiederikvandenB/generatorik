package generate

import (
  "log"
  "text/template"
)

type StoriesProps struct {
  ComponentName string
}

func WriteStories(props StoriesProps, fileName string) {
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
