package main

import (
	"buildurl"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"

	"github.com/tidwall/gjson"
)

func main() {

	//	APIURL := "https://gitea-tooling.az.globaldevops.gdpdentsu.net/api/v1/repos/search?q=Joe&access_token=b479bc9baa5ae757a458d73b2a610eafce6e4d5a"	//fmt.Println(FullURL())
	APIURL := buildurl.FullURL().String() //"https://gitea-tooling.az.globaldevops.gdpdentsu.net/api/v1/repos/search?q=Joe&access_token=b479bc9baa5ae757a458d73b2a610eafce6e4d5a"
	//fmt.Println(buildurl.FullURL().String())
	req, err := http.NewRequest(http.MethodGet, APIURL, nil)
	if err != nil {
		panic(err)
	}
	client := http.DefaultClient
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	//fmt.Printf(string(body))
	value := gjson.Get(string(body), "data.#.clone_url")
	//string avalue = strings.Replace(avalue, "[]", "", -1 )
	//println(value.String())
	if !value.Exists() {
		fmt.Println("No Repository called")
		os.Exit(1)
	} else {
		url1 := value.String()
		url2 := strings.Replace(url1, "[", "", -1)
		finalUrl := strings.Replace(url2, "]", "", -1)
		fmt.Println(finalUrl)
	}

	//fmt.Printf("%v", string(value)

	//fmt.Printf("%v", string(body))
}
