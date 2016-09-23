package main

import "bytes"
import "net/http"
import "fmt"
import "io/ioutil"


func PostNewJenkins() {
    fmt.Println("Making post call to create a job in jenkins");
    client := &http.Client{}
    jenkinsURL :="http://ekart-jenkins-master.nm.flipkart.com:8080"
    jobName := "jobFromGoLang4"
    url := jenkinsURL+"/createItem?name="+jobName
    fmt.Println("URL:>", url)

    dat, err := ioutil.ReadFile("/Users/ambarish.a/Documents/myRepos/GoLangSamples/src/github.com/ambsflip/jenkinsGo/jenkinsJobTemplate.xml")
    check(err)
    jenkinsTemplate :=string(dat)
    body := jenkinsTemplate
    fmt.Println("body :",body)
    //body := "<flow-definition plugin=\"workflow-job@2.2\"><actions/><description/><keepDependencies>false</keepDependencies><properties><com.synopsys.arc.jenkins.plugins.ownership.jobs.JobOwnerJobProperty plugin=\"ownership@0.8\"/></properties><definition class=\"org.jenkinsci.plugins.workflow.cps.CpsScmFlowDefinition\" plugin=\"workflow-cps@2.4\"><scm class=\"hudson.plugins.git.GitSCM\" plugin=\"git@2.4.1\"><configVersion>2</configVersion><userRemoteConfigs><hudson.plugins.git.UserRemoteConfig><url>git@github.com:Flipkart/GenesisSampleApp.git</url></hudson.plugins.git.UserRemoteConfig></userRemoteConfigs><branches><hudson.plugins.git.BranchSpec><name>*/master</name></hudson.plugins.git.BranchSpec></branches><doGenerateSubmoduleConfigurations>false</doGenerateSubmoduleConfigurations><submoduleCfg class=\"list\"/><extensions/></scm><scriptPath>Jenkinsfile</scriptPath></definition><triggers/></flow-definition>"
    //var jsonStr = []byte(`{"tokens":[{"name":"javaPackagePath","value":"com/ekart/springbootstarter","type":"path"},{"name":"javaPackageName","value":"com.ekart.springbootstarter","type":"content"},{"name":"javaPackagePathName","value":"com/ekart/springbootstarter","type":"content"},{"name":"zipkinApplicationName","value":"spring-boot-jetty-zipkin","type":"content"},{"name":"projectName","value":"spring-boot-jetty-starter","type":"content"},{"name":"teamEmail","value":"team-alias@flipkart.com","type":"content"},{"name":"debPackageDescription","value":"ekart spring boot jetty package","type":"content"},{"name":"deploymentUserName","value":"fk-deployment-user-name","type":"content"},{"name":"deploymentUserId","value":"121212","type":"content"},{"name":"deploymentUserGroup","value":"fk-deployment-user-group","type":"content"},{"name":"deploymentUserGroupId","value":"343434","type":"content"},{"name":"swaggerText","value":"Spring Boot + Jetty Starter template","type":"content"},{"name":"jenkinsUrl","value":"http://localhost:8080","type":"jenkins"},{"name":"jenkinsJobName","value":"a","type":"jenkins"},{"name":"pipelineScriptGitUrl","value":"https://github.com/Flipkart/ekart-starter-kit.git","type":"jenkins"},{"name":"gitBranchName","value":"master","type":"jenkins"},{"name":"pipelineScriptFileName","value":"Jenkinsfile","type":"jenkins"}],"inputs":[{"name":"templateKey","value":"1"},{"name":"gitUrl","value":""},{"name":"username","value":""},{"name":"password","value":""}]}`)
    req, err := http.NewRequest("POST", url, bytes.NewBuffer([]byte(body)))
    //req.Header.Set("Content-Length", "123")
    req.Header.Set("Content-Type", "text/xml")
    req.SetBasicAuth("ambarish.a", "")

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

func PostNewJenkins2(){

    body := "<flow-definition plugin=\"workflow-job@2.2\"><actions/><description/><keepDependencies>false</keepDependencies><properties><com.synopsys.arc.jenkins.plugins.ownership.jobs.JobOwnerJobProperty plugin=\"ownership@0.8\"/></properties><definition class=\"org.jenkinsci.plugins.workflow.cps.CpsScmFlowDefinition\" plugin=\"workflow-cps@2.4\"><scm class=\"hudson.plugins.git.GitSCM\" plugin=\"git@2.4.1\"><configVersion>2</configVersion><userRemoteConfigs><hudson.plugins.git.UserRemoteConfig><url>git@github.com:Flipkart/GenesisSampleApp.git</url></hudson.plugins.git.UserRemoteConfig></userRemoteConfigs><branches><hudson.plugins.git.BranchSpec><name>*/master</name></hudson.plugins.git.BranchSpec></branches><doGenerateSubmoduleConfigurations>false</doGenerateSubmoduleConfigurations><submoduleCfg class=\"list\"/><extensions/></scm><scriptPath>Jenkinsfile</scriptPath></definition><triggers/></flow-definition>"
    client := &http.Client{}
    url := "http://ekart-jenkins-master.nm.flipkart.com:8080/createItem?name=JobFromGO"
    // build a new request, but not doing the POST yet
    req, err := http.NewRequest("POST",url, bytes.NewBuffer([]byte(body)))
    if err != nil {
        fmt.Println(err)
    }
    // you can then set the Header here
    // I think the content-type should be "application/xml" like json...
    req.Header.Add("Content-Type", "application/xml; charset=utf-8")
    req.SetBasicAuth("ambarish.a", "S@iram1986")
    // now POST it
    resp, err := client.Do(req)
    if err != nil {
        fmt.Println(err)
    }
    defer resp.Body.Close()
    fmt.Println(resp)

}
