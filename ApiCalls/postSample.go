package apicall

import "bytes"
import "net/http"
import "fmt"
import "io/ioutil"
import "reflect"

func postCall() {
    fmt.Println("Making post call to generate a project in genesis");
    url := "http://localhost:28028/generators/generateTemplate"
    fmt.Println("URL:>", url)

    var jsonStr = []byte(`ENTER A JSON STRING`)
    req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonStr))
    req.Header.Set("Content-Length", "123")
    req.Header.Set("Content-Type", "application/json")
    //req.SetBasicAuth(username, password)
    client := &http.Client{}
    resp, err := client.Do(req)
    if err != nil {
        panic(err)
    }
    defer resp.Body.Close()

    fmt.Println("response Status:", resp.Status)
    fmt.Println("resp.StatusCode:",resp.StatusCode, reflect.TypeOf(resp.StatusCode))
    fmt.Println("response Header:", resp.Header, reflect.TypeOf(resp.Header))
    fmt.Println("response header.Get", resp.Header.Get("X-Error"))
    fmt.Println("response Headers:", resp.Header)
    body, _ := ioutil.ReadAll(resp.Body)
    fmt.Println("response Body:", string(body))
}
