package main

import (
	//"buildurl"

	//"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	//"strings"

	"github.com/tidwall/gjson"
	//"strings"
	//"github.com/tidwall/gjson"
	//"os"
	//"strings"
)

type Username struct {
	Username string `json:"username"`
}

type OrgName struct {
	OrgString string
}

func main() {

	APIURL := "https://gitea-tooling.az.globaldevops.gdpdentsu.net/api/v1/admin/orgs?access_token=b479bc9baa5ae757a458d73b2a610eafce6e4d5a" //fmt.Println(FullURL())
	//APIURL := buildurl.FullURL().String() //"https://gitea-tooling.az.globaldevops.gdpdentsu.net/api/v1/repos/search?q=Joe&access_token=b479bc9baa5ae757a458d73b2a610eafce6e4d5a"
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

	str := fmt.Sprintf("%s", body)
	value := gjson.Get(str, "#.username")

	for _, newvalue := range value.Array() {
		fmt.Println(newvalue.String())
	}

}
