package main

import (
   "fmt"
   "os"
   "path/filepath"
)

func main (){
  fmt.Println(ListFileInPwd("jar"))
}
func ListFileInPwd(Extension string) (string){

   dirname := "." + string(filepath.Separator)
   d, err := os.Open(dirname)
   if err != nil {
       fmt.Println(err)
       os.Exit(1)
   }
   defer d.Close()

   files, err := d.Readdir(-1)
   if err != nil {
       fmt.Println(err)
       os.Exit(1)
   }

   //fmt.Println("Reading "+ dirname)
   Ext := "."+Extension
   for _, file := range files {
       if file.Mode().IsRegular() {
           if filepath.Ext(file.Name()) == Ext {
             return file.Name()
           }
       }
   }
   return ""
}
