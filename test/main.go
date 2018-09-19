package main

import (
  "github.com/fluent/fluent-logger-golang/fluent"
  "fmt"
  //"time"
)

func main() {
  logger, err := fluent.New(fluent.Config{
    FluentHost: "127.0.0.1",
    FluentPort: 24284,
  })
  if err != nil {
    fmt.Println(err)
  }
  defer logger.Close()
  tag := "debug.export.office"
  var data = map[string]string{
    "foo":  "bar",
    "hoge": "hoge",
  }
  error := logger.Post(tag, data)
  // error := logger.PostWithTime(tag, time.Now(), data)
  if error != nil {
    panic(error)
  }
}