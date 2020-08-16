package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/tidwall/gjson"
)

type Result struct {
	gjson.Result
}

func main() {

	APIURL := "https://api.github.com/user/orgs?access_token=e459fbb080b05bd3ba6a52220c06a8e46c7fac77"
	//APIURL := buildurl.FullURL().String() //"https://gitea-tooling.az.globaldevops.gdpdentsu.net/api/v1/repos/search?q=Joe&access_token=b479bc9baa5ae757a458d73b2a610eafce6e4d5a"

	req, err := http.NewRequest(http.MethodGet, APIURL, nil)
	if err != nil {
		panic(err)
	}
	req.Header.Add("content-type", "application/json")

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
	//fmt.Println(body)

	aString := string(body)
	//fmt.Println(s)
	//fmt.Println(reflect.TypeOf(s))
	value := gjson.Get(string(aString), "#.login")
	//string avalue = strings.Replace(avalue, "[]", "", -1 )

	if !value.Exists() {
		fmt.Println("No Orginsation exists")
		os.Exit(1)
	} else {
		fmt.Println(value)
	}

	//fmt.Printf("%v", string(value)

	//fmt.Printf("%v", string(body))
}

func gitea() {

}
