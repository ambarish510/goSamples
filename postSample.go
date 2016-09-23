package main

import "bytes"
import "net/http"
import "fmt"
import "io/ioutil"
import "reflect"

func postCall() {
    fmt.Println("Making post call to generate a project in genesis");
    url := "http://localhost:28028/generators/generateTemplate"
    fmt.Println("URL:>", url)

    var jsonStr = []byte(`{"tokens":[{"name":"javaPackagePath","value":"com/ekart/springbootstarter","type":"path"},{"name":"javaPackageName","value":"com.ekart.springbootstarter","type":"content"},{"name":"javaPackagePathName","value":"com/ekart/springbootstarter","type":"content"},{"name":"zipkinApplicationName","value":"spring-boot-jetty-zipkin","type":"content"},{"name":"projectName","value":"spring-boot-jetty-starter","type":"content"},{"name":"teamEmail","value":"team-alias@flipkart.com","type":"content"},{"name":"debPackageDescription","value":"ekart spring boot jetty package","type":"content"},{"name":"deploymentUserName","value":"fk-deployment-user-name","type":"content"},{"name":"deploymentUserId","value":"121212","type":"content"},{"name":"deploymentUserGroup","value":"fk-deployment-user-group","type":"content"},{"name":"deploymentUserGroupId","value":"343434","type":"content"},{"name":"swaggerText","value":"Spring Boot + Jetty Starter template","type":"content"},{"name":"jenkinsUrl","value":"http://localhost:8080","type":"jenkins"},{"name":"jenkinsJobName","value":"a","type":"jenkins"},{"name":"pipelineScriptGitUrl","value":"https://github.com/Flipkart/ekart-starter-kit.git","type":"jenkins"},{"name":"gitBranchName","value":"master","type":"jenkins"},{"name":"pipelineScriptFileName","value":"Jenkinsfile","type":"jenkins"}],"inputs":[{"name":"templateKey","value":"1"},{"name":"gitUrl","value":""},{"name":"username","value":""},{"name":"password","value":""}]}`)
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
