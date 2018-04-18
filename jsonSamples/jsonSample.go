package main

import "encoding/json"
import "fmt"
import "io/ioutil"
import "log"
import "net/http"
import "time"
import "os"

func main() {
  jsonprocessing()

}


type Response2 struct {
    Page   int      `json:"page"`
    Fruits []string `json:"fruits"`
}

func jsonprocessing(){
 //Marshal is to convert into json
 //UnMarshal is to convert json into other types

 //create a nested json
 //create a array of json

    res2D := &Response2{
      Page:   1,
      Fruits: []string{"apple", "peach", "pear"}}
    res2B, _ := json.Marshal(res2D)
    fmt.Println(string(res2B))

    byt := []byte(`{"num":6.13,"strs":["a","b"]}`)
    // Reading json to general map string : interface
    var dat map[string]interface{}
    if err := json.Unmarshal(byt, &dat); err != nil {
        panic(err)
    }
    fmt.Println(dat)
    num := dat["num"].(float64)
    fmt.Println(num)
    strs := dat["strs"].([]interface{})
    str1 := strs[0].(string)
    fmt.Println(str1)

    //Reading the json into a struct
    str := `{"page": 1, "fruits": ["apple", "peach"]}`
    res := Response2{}
    json.Unmarshal([]byte(str), &res)
    fmt.Println(res)
    fmt.Println(res.Fruits[0])

    str2 := `[{"artifact":{"package_name":"com.github.tomakehurst:wiremock-standalone","version":"2.6.0"},"status":"DOWNLOAD_FAILURE","store":"MAVEN","reason":"Error in resolving artifact{\"package_name\":\"com.github.tomakehurst:wiremock-standalone\",\"version\":\"2.6.0\"}with message: Could not transfer artifact com.github.tomakehurst:wiremock-standalone:pom:javadocs:2.6.0 from/to artifactory-libs-snapshot-local (http://artifactory.nm.flipkart.com:8081/artifactory/libs-snapshot-local/): Conflict (409)","external_download_url":"http://repo1.maven.org/maven2/"}]`
    res2 := []SearchResponse{}
    json.Unmarshal([]byte(str2), &res2)
    //fmt.Println(res)
    fmt.Println(res2[0].Store)

}


type SearchResponse struct {
  // ArtifactDetails struct {
  //   package_name string `json:"package_name"`
  //   version string `json:"version"`
  // }
  Store string `json:"store"`
	Status string `json:"status"`
  Reason string `json:"reason"`
  ExtDwnldURL string `json:"external_download_url"`
}
type SearchResponses struct {
  searchresponses []SearchResponse
}

func jsonProcessing3() {
  //json to Array , no nested structure involved

	body := []byte(`["status":"DOWNLOAD_FAILURE","store":"MAVEN","reason":"Error in resolving artifact{\"package_name\":\"com.github.tomakehurst:wiremock-standalone\",\"version\":\"2.6.0\"}with message: Could not transfer artifact com.github.tomakehurst:wiremock-standalone:pom:javadocs:2.6.0 from/to artifactory-libs-snapshot-local (http://artifactory.nm.flipkart.com:8081/artifactory/libs-snapshot-local/): Conflict (409)","external_download_url":"http://repo1.maven.org/maven2/"}]`)
	keys := make([]SearchResponse,0)
	json.Unmarshal(body, &keys)

  fmt.Println("Store",keys[0].Store)
  fmt.Println("Status",keys[0].Status)
  fmt.Println("Reason",keys[0].Reason)
  fmt.Println("ExtDwnldURL",keys[0].ExtDwnldURL)
}

type Astros struct{
  Number int `json:"number"`
  Message string `json:"message"`
  People []Person `json:"people"`
}

type Person struct {
  Name string `json:"name"`
  Craft string `json:"craft"`

}
// read json api response into
func jsonProcessing1() {
    url := "http://api.open-notify.org/astros.json"

    httpClient := http.Client{
        Timeout: time.Second * 2, // Maximum of 2 secs
    }
    req, err := http.NewRequest(http.MethodGet, url, nil)
    if err != nil {
        log.Fatal(err)
    }
    res, getErr := httpClient.Do(req)
    if getErr != nil {
        log.Fatal(getErr)
    }
    body, readErr := ioutil.ReadAll(res.Body)
    if readErr != nil {
        log.Fatal(readErr)
    }

    var dat map[string]interface{}
    jsonErr := json.Unmarshal(body, &dat)
    if jsonErr != nil {
        log.Fatal(jsonErr)
    }
    fmt.Println(dat)
    num := dat["number"].(float64)
    fmt.Println(num)
    msg := dat["message"].(string)
    fmt.Println(msg)
    peopleMap := dat["people"].([]interface{})
    fmt.Println(peopleMap)

    // Reading to struct with partial fields
    astro := Astros{}
    json.Unmarshal(body, &astro)
    fmt.Println(astro)
    fmt.Println(astro.Number)
    fmt.Println(astro.Message)
    fmt.Println(astro.People[0])
    fmt.Println(astro.People[1].Craft)

}

type ReleaseDetails struct{
   TagName string `json:"tag_name"`
   PreRelease bool `json:"prerelease"`
   Assets []AssetDetails `json:"assets"`
}

type AssetDetails struct{
  AssetURL string `json:"url"`
  AssetName string `json:"name"`
  ContentType string `json:"content_type"`
  Size int `json:"size"`
}

func githubLatestRelease() {
    url := "https://api.github.com/repos/Flipkart/artifactory-service/releases/latest?access_token=6bd12c8d5ef9a1b9a39478e812a0e231717b53e1"

    httpClient := http.Client{
        Timeout: time.Second * 120, // Maximum of 2 minutes
    }
    req, err := http.NewRequest(http.MethodGet, url, nil)
    if err != nil {
        log.Fatal(err)
    }
    res, getErr := httpClient.Do(req)
    if getErr != nil {
        log.Fatal(getErr)
    }
    body, readErr := ioutil.ReadAll(res.Body)
    if readErr != nil {
        log.Fatal(readErr)
    }

    var dat map[string]interface{}
    jsonErr := json.Unmarshal(body, &dat)
    if jsonErr != nil {
        log.Fatal(jsonErr)
    }
    //fmt.Println(dat)
    assetUrl := dat["url"].(string)
    fmt.Println(assetUrl)
    tag_name := dat["tag_name"].(string)
    fmt.Println(tag_name)
    //authorDetails := dat["author"].(interface{})
    //fmt.Println(authorDetails)
    assetDetails := dat["assets"].([]interface{})
    fmt.Println(assetDetails[0])
    //fmt.Println(assetDetails[0]["url"].(string))


    // // Reading to struct with partial fields
    release := ReleaseDetails{}
    json.Unmarshal(body, &release)
    //fmt.Println(astro)
    fmt.Println(release.TagName)
    fmt.Println(release.PreRelease)
    fmt.Println(release.Assets[0].AssetURL)
    fmt.Println(release.Assets[0].AssetName)
    fmt.Println(release.Assets[0].ContentType)
    fmt.Println(release.Assets[0].Size)
}
func jsonProcessing2(){

  byt := []byte(`{"num":6.13,"strs":["a","b"]}`)
  var dat map[string]interface{}
  if err := json.Unmarshal(byt, &dat); err != nil {
        panic(err)
    }
  fmt.Println("dat",dat)
  num := dat["num"].(float64)
  fmt.Println("num",num)

  strs := dat["strs"].([]interface{})
  str1 := strs[0].(string)
  fmt.Println("str1",str1)

  str := `{"page": 1, "fruits": ["apple", "peach"]}`
  res := Response2{}
  json.Unmarshal([]byte(str), &res)
  fmt.Println("res",res)
  fmt.Println("res.Fruits[0]",res.Fruits[0])

  enc := json.NewEncoder(os.Stdout)
  d := map[string]int{"apple": 5, "lettuce": 7}
  enc.Encode(d)
}
