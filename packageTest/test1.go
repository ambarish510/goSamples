package PackageTest

import(
    "fmt"
)
//witin the same package, fucntions in one file can refer to variables in other fileDescriptor
// this can be used without qualifiers before the variable name.
func GetConfigValues() {
  LogFileName = "_artLogFile"
  LogFilePath = "logs"
  TimeoutForHttpRequest = 5

  //return logConfig,nil
}

func PrintValues(){
  fmt.Println(LogFileName, LogFilePath)
}
