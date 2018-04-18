package main

//import "bytes"
import "net/http"
import "time"
//import "runtime"
import "os/exec"
import "fmt"
import "io"
import "os"
//import "reflect"
import "github.com/inconshreveable/go-update"
import "github.com/Flipkart/artifactory-service/client/artcli/utils"

func main() {
  //m := make(map[string]string)
  //m["Accept"] = "application/json"
  //fmt.Println("m  ",m)
  //resp,err := DownloadFromGithub("https://api.github.com/repos/stedolan/jq/commits?per_page=5","application/json")
  //defer resp.Body.Close()
  //fmt.Println("resp  ",resp,"\nerr ", err)
  test2()
  //test1()
  //ExampleCmd_Start()
}


func test2() {
    AssetURL :="https://api.github.com/repos/Flipkart/buildinfra_jenkins_pipeline_demo_cloud/releases/assets/4311562?access_token=6bd12c8d5ef9a1b9a39478e812a0e231717b53e1"
    TimeOut := 600
    //resp,err := DownloadFromGithub(TimeOut,AssetURL,"application/octet-stream")
    resp,err := utils.MakeGetCall(TimeOut,AssetURL,"application/octet-stream")
    fmt.Println("err :",err)
    filepath:="/Users/ambarish.a/Documents/myRepos/GoLangSamples/src/github.com/ambsflip/goSamples/ApiCalls/SampleJarFile.jar"

    Err:=utils.WriteResonseToDisk(resp,filepath)
    fmt.Println("Err:",Err)
}

func DownloadFromGithub(TimeOut int,getURL string, contentType string)(*http.Response, error){
  fmt.Println("getURL",getURL)
  fmt.Println("contentType",contentType)
  client := &http.Client{
    Timeout:  time.Duration(TimeOut) * time.Second,
  }
  req, err := http.NewRequest("GET", getURL, nil)
  req.Header.Add("Accept", contentType)
  resp,err := client.Do(req)

  return resp,err

}
func writeResonseToDisk(resp *http.Response, destPath string) error {
	out, err := os.Create(destPath)
	if err != nil  {
		return err
	}

	defer out.Close()
	defer resp.Body.Close()

	_, err = io.Copy(out, resp.Body)
  err2 := os.Chmod(destPath, 0777)
  if err2 != nil {
     fmt.Println(err)
  }

	return err
}
func ExampleCmd_Start() {
  //curl -LJOH Accept: application/octet-stream https://api.github.com/repos/Flipkart/buildinfra_jenkins_pipeline_demo_cloud/releases/assets/4296294?access_token=6bd12c8d5ef9a1b9a39478e812a0e231717b53e1
  	err := exec.Command("curl", "-L","-J","-O","-H", "'Accept: application/octet-stream'", "https://api.github.com/repos/Flipkart/artifactory-service/releases/assets/4167867?access_token=6bd12c8d5ef9a1b9a39478e812a0e231717b53e1").Run()
    //err := cmd.Start()
    if err != nil {
  		fmt.Println(err)
  	}
    fmt.Println("completed")
}

func test1() {
    client := &http.Client{
      Timeout:  time.Duration(60) * time.Second,
      //CheckRedirect: nil,
    }

    req, err := http.NewRequest("GET","https://api.github.com/repos/Flipkart/artifactory-service/releases/assets/4167867?access_token=6bd12c8d5ef9a1b9a39478e812a0e231717b53e1",nil)
    req.Header.Add("Accept", "application/octet-stream")
    resp,err := client.Do(req)
    if err != nil {
        fmt.Printf("http.Get => %v", err.Error())
    }
    defer resp.Body.Close()
    // client tried to access.
    finalURL := resp.Request.URL.String()

    fmt.Printf("The URL you ended up at is: %v\n", finalURL)
    //fmt.Println("Body:\n",resp.Body)

    //this is to update the invoked binay with new data
    err = update.Apply(resp.Body, update.Options{})
    if err != nil {
      fmt.Println(err)
      return
    }

    // // Create the file
    // filepath:="/Users/ambarish.a/Documents/myRepos/GoLangSamples/src/github.com/ambsflip/goSamples/ApiCalls/NewArt"
    // out, err := os.Create(filepath)
    // if err != nil  {
    //   fmt.Println(err)
    //   return
    // }
    // defer out.Close()
    //
    // // Write the body to file
    // _, err = io.Copy(out, resp.Body)
    // if err != nil  {
    //   fmt.Println(err)
    //   return
    // }



}

func MakeGetCall(getURL string, headers map[string]string) (*http.Response, error) {
  client := &http.Client{
    Timeout:  time.Duration(5) * time.Second,
  }
  req, err := http.NewRequest("GET", getURL, nil)
  for k, v := range headers {
    fmt.Println("k:", k, "v:", v)
    req.Header.Add(k, v)
  }

  resp,err := client.Do(req)

  return resp,err
}
