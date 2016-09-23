package main

import (
    "fmt"
    "encoding/json"
    "gopkg.in/yaml.v2"
)
type Person struct {
    Name string `json:"name" yaml:"personname"` // Use "name" instead of "Name" for the JSON
    Age int `json:"age" yaml:"personage"`      // field name.
}
func jsonProcessing(){
  // Create a Person struct.
  p := Person{"John", 30}
  // Convert it to JSON.
  j, err := json.Marshal(p)
  if err != nil {
    fmt.Println(err)
  }
  fmt.Println(string(j))

  byt := []byte (`{"name":"Johnny","age":31}`)
  var p2 Person
  err2 := json.Unmarshal(byt, &p2)
  if err2 != nil {
       fmt.Println( err2)
  }
  fmt.Println(p2.Name,"\t\t",p2.Age)
}

func yamlProcessing(){
  p := Person{"Sachin", 40}
  // Convert the Person struct to YAML.
  y, err := yaml.Marshal(p)
  if err != nil {
    fmt.Println(err)
  }
  fmt.Println(string(y))

  byt := []byte (`{"personname":"Dravid","personage":42}`)
  var p2 Person
  err2 := yaml.Unmarshal(byt, &p2)
  if err2 != nil {
       fmt.Println( err2)
  }
  fmt.Println(p2.Name,"\t\t",p2.Age)
}
