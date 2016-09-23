package main

import (
  "os"
  "fmt"
  "log"
)

/*CreateLogFile creates a file named testlogfile in the PWD/log directory and
writes a log entry to it. If the file already exists, the new log entry is
appended to it. if the log directory doesnt exists it throws an error.
*/
func CreateLogFile(){
  f, err := os.OpenFile("logs/testlogfile", os.O_RDWR | os.O_CREATE | os.O_APPEND, 0666)
  if err != nil {
    fmt.Println("error opening file: %v", err)
  }
  defer f.Close()

  log.SetOutput(f)
  log.Println("This is a test log entry")
}
