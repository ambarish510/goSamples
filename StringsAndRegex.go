package main

import "fmt"
import "regexp"
import "net/url"
import "strings"
import "strconv"

func regexsample() {

  body := "<flow-definition plugin=\"workflow-job@2.2\"><actions/><description/><keepDependencies>false</keepDependencies><properties><com.synopsys.arc.jenkins.plugins.ownership.jobs.JobOwnerJobProperty plugin=\"ownership@0.8\"/></properties><definition class=\"org.jenkinsci.plugins.workflow.cps.CpsScmFlowDefinition\" plugin=\"workflow-cps@2.4\"><scm class=\"hudson.plugins.git.GitSCM\" plugin=\"git@2.4.1\"><configVersion>2</configVersion><userRemoteConfigs><hudson.plugins.git.UserRemoteConfig><url>git@github.com:Flipkart/GenesisSampleApp.git</url></hudson.plugins.git.UserRemoteConfig></userRemoteConfigs><branches><hudson.plugins.git.BranchSpec><name>*/master</name></hudson.plugins.git.BranchSpec></branches><doGenerateSubmoduleConfigurations>false</doGenerateSubmoduleConfigurations><submoduleCfg class=\"list\"/><extensions/></scm><scriptPath>Jenkinsfile</scriptPath></definition><triggers/></flow-definition>"

  r, _ := regexp.Compile("<hudson.plugins.git.UserRemoteConfig><url>.+</url>")
  fmt.Println(r.ReplaceAllString(body, "<hudson.plugins.git.UserRemoteConfig><url>GIT URL</url>"))
}
func concatenateStrings(){
  str1 := "hello"
  str2 := "world"
  str3 := str1+" "+str2
  fmt.Println(str3)
}
func URLFormatCheck(){
    s := "git@github.com:ambsFlip/stringutil.git"
    u, err := url.Parse(s)
    if err != nil {
        panic(err)
    }
  fmt.Println("Scheme",u.Scheme)
  fmt.Println("Host", u.Host)
}
func StringReplace(){
  responseStatus := "401 Invalid password/token for user: ambarish.a"
  respStatusCode:= 401

  reponseStatusMsg := strings.Replace(responseStatus, strconv.Itoa(respStatusCode), "", -1)
  fmt.Println("responseStatus",responseStatus)
  fmt.Println("respStatusCode",strconv.Itoa(respStatusCode))
  fmt.Println("reponseStatusMsg",reponseStatusMsg)
}
