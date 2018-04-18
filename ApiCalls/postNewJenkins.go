package apicall

import "bytes"
import "net/http"
import "fmt"
import "io/ioutil"


func PostNewJenkins() {
    fmt.Println("Making post call to create a job in jenkins");
    client := &http.Client{}
    jenkinsURL :="jenkinsURL"
    jobName := "jobFromGoLang4"
    url := jenkinsURL+"/createItem?name="+jobName
    fmt.Println("URL:>", url)

    dat, err := ioutil.ReadFile("/Users/ambarish.a/Documents/myRepos/GoLangSamples/src/github.com/ambsflip/jenkinsGo/jenkinsJobTemplate.xml")
    //check(err)
    jenkinsTemplate :=string(dat)
    body := jenkinsTemplate
    fmt.Println("body :",body)
    req, err := http.NewRequest("POST", url, bytes.NewBuffer([]byte(body)))
    //req.Header.Set("Content-Length", "123")
    req.Header.Set("Content-Type", "text/xml")
    req.SetBasicAuth("username", "password")

    resp, err := client.Do(req)
    if err != nil {
        panic(err)
    }
    defer resp.Body.Close()

    fmt.Println("response Status:", resp.Status)
    fmt.Println("response Headers:", resp.Header)
    //body, _ := ioutil.ReadAll(resp.Body)
    fmt.Println("response Body:", resp.Body)
}
