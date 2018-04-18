package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func main() {
	gopath := "/Users/ambarish.a/Documents/myRepos/GoLangSamples/src/github.com/ambsflip/goSamples/FileAndDir"
	getShellScript(gopath,".job")
}

func getShellScript(dirpath string,ext string) {
	err := filepath.Walk(dirpath, func(path string, info os.FileInfo, err error) error {
		if info.IsDir() {
			return nil
		}
		if filepath.Ext(path) == ext {
			fmt.Println(path)
			err = os.Remove(path)
		}
		return nil
	})
	if err != nil {
		fmt.Printf("Error in removing the old jars [%v]\n", err)
	}
}
