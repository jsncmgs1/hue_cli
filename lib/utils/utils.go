package utils

import (
  "fmt"
  "bytes"
  "encoding/json"
)

func PrettyPrint(jsonBytes []byte){
  buf := new(bytes.Buffer)
  json.Indent(buf, jsonBytes, "", "  ")
  fmt.Println(buf)
}
