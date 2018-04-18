package main

import (
  "bytes"
  "fmt"
  "io"
  //"log"
  "mime/multipart"
  "net/http"
  "os"
  "path/filepath"
)

// Creates a new file upload http request with optional extra params
func newfileUploadRequest(uri string, paramName, path string) (*http.Response, error) {
  file, err := os.Open(path)
  if err != nil {
      return nil, err
  }
  defer file.Close()

  body := &bytes.Buffer{}
  writer := multipart.NewWriter(body)
  part, err := writer.CreateFormFile(paramName, filepath.Base(path))
  if err != nil {
      return nil, err
  }
  _, err = io.Copy(part, file)

  err = writer.Close()
  if err != nil {
      return nil, err
  }

  req, err := http.NewRequest("POST", uri, body)
  req.Header.Set("Content-Type", writer.FormDataContentType())


  if err != nil {
      //log.Fatal(err)
      fmt.Println(err)
  }
  client := &http.Client{}
  resp, err := client.Do(req)

  return resp, err
}

func main() {
  // path, _ := os.Getwd()
  // path += "/test.pdf"
  resp, err := newfileUploadRequest("http://10.85.58.239/artifactory/v1.0/artifacts/upload/pypi/59f198d1b9b2bf14ddf20ef7/thespian-3.8.1.zip", "file", "/Users/ambarish.a/Downloads/thespian_3.8.1.zip")
  if err != nil {
      //log.Fatal(err)
      fmt.Println(err)
  } else {
      body := &bytes.Buffer{}
      _, err := body.ReadFrom(resp.Body)
    if err != nil {
          //log.Fatal(err)
          fmt.Println(err)
      }
    resp.Body.Close()
      fmt.Println(resp.StatusCode)
      fmt.Println(resp.Header)
      fmt.Println(body)
  }
}
