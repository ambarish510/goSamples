package main

import (
  "fmt"
  )

/*
go get -u github.com/jteeuwen/go-bindata/...
go-bindata -o myfile.go data/config.yaml

this will create the go equivalent file (myfile.go) for the files inside the
directory data

now to access the data items in the config files, include the myfile.go into
required package and use the following functions
*/

func convertTextToGo(){
  content, err := Asset("data/config.yaml")
  if err!=nil{
    fmt.Println("unable to read config file")
  }else{
    configContent := string (content)
    fmt.Println(configContent)
  }
}
