package main
import (
    //"bufio"
    "fmt"
    //"io"
    "io/ioutil"
    "os"
    "os/user"
)

func check(e error) {
    if e != nil {
        panic(e)
    }
}
func readFile(){
    dat, err := ioutil.ReadFile("/Users/ambarish.a/Documents/myRepos/GoLangSamples/src/github.com/ambsflip/goSamples/jenkinsJobTemplate.xml")
    if err != nil {
        fmt.Println(err)
    }
    beforeReplacement :=string(dat)
    fmt.Print("beforeReplacement :", beforeReplacement)
}

func getWorkingDir(){
  path,_ := os.Getwd()
  fmt.Println("Working Directory",path)
  usr,_ := user.Current()
  fmt.Println( usr.HomeDir )
}

//function isFileExists can be used to check the existance of directories also
func isFileExists(path string) (bool, error) {
    _, err := os.Stat(path)
    if err == nil { return true, nil }
    if os.IsNotExist(err) { return false, nil }
    return true, err
}
