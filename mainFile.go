package main
import (
  "fmt"
  "github.com/ambsflip/goSamples/MyMath"
)
var(
  Version string
  Build string
)


func main() {
  //cliSample()
  //fmt.Println("Exiting")
  //fmt.Println("Version, Build : ",Version,Build)

  //fmt.Println(average(98,93,77,81,84))

  //CONCURRENCY
  // for i := 0; i < 5; i++ {
  //     go f(i)
  //   }
  // var input string
  // fmt.Scanln(&input)
  // fmt.Println("input is :",input)

  // CHANNELS
  // var c chan string = make(chan string)
  // go pinger(c)
  // go ponger(c)
  // go printer(c)
  // var input string
  // fmt.Scanln(&input)

  //package - note the use of uppercase first alphabet for imported function
  fmt.Println(MyMath.AverageOf(100,98,93,77,81,84))

  //fmt.Println(os.Args)
  //cmdLineSample()
  //cmdLineFlags()
  //postCall()
  //PostNewJenkins()
  //readFile()
  //regexsample();
  //readCredentials();
  //concatenateStrings();
  //URLFormatCheck()
  //StringReplace();
  //getWorkingDir()
  // fileExists,errFileExists := isFileExists("/Users/ambarish.a/Documents/myRepos/GoLangSamples/src/github.com/ambsflip/goSamples")
  // if errFileExists != nil {
  //     fmt.Println("errFileExists: ",errFileExists)
  // }
  // fmt.Println("fileExists: ",fileExists)
  // CreateLogFile()
  // convertTextToGo()
  // jsonProcessing()
  // yamlProcessing()
}
