package main

import "fmt"
import "os/exec"



func main() {
  //directoryPath:= "/Users/ambarish.a/Documents/testjar/"
  //JarPath := directoryPath+"Sample.jar"
  //Params := "Bob"
  PomPath := "/Users/ambarish.a/Documents/testjar/pom.xml"
  JarPath := "dependency-resolver-1.0-SNAPSHOT.jar"
  Params := "com.flipkart.depcheck.app -buildFile "+PomPath+" -create"
  // fmt.Println("JarPath",JarPath)
  // //cmd := exec.Command(java, "-jar",JarPath , "-buildFile", PomPath, "-create")
  // cmd := exec.Command("java","-jar",JarPath)
  // err := cmd.Run()
  //
  // if err != nil {
	// 	fmt.Println("an error occurred.\n")
	// 	fmt.Println(err)
	// }
  RunJar(JarPath+" "+Params)
}
//Function that takes the jar path and its params as a single string
//It will execute the jar with the given params and the output is
//printed to stdout. It catches the errors (if any) and prints it to
// stdout. There is no return value.
func RunJar(jarPathAndParams string) {

    fmt.Println(jarPathAndParams," ")
    cmd_prep := "java -Xmx2g -jar "+jarPathAndParams
    fmt.Println("cmd_prep",cmd_prep)
    cmd_output, err := exec.Command("bash", "-c", cmd_prep).Output()
    if err != nil {
        fmt.Println(err)
    }
    fmt.Println(string(cmd_output))
}
