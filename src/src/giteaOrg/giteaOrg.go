package main

import(
	"fmt"
	"io/ioutil"
	"net/http"
//	"strings"
)
func main() {

	//APIURL := buildurl.FullURL().String() 
	//url := fmt.Sprintf(giteaapiurl, project) + "org/" + orgname
	//token := "b479bc9baa5ae757a458d73b2a610eafce6e4d5a"
	token := "219cba8a7cf6249009cffaa11cef2424432b9bae"
	APIURL := "https://gitea-tooling.az.transformers.gdpdentsu.net/api/v1/admin/orgs?page=1&limit=10"
	//fmt.Println(buildurl.FullURL().String())
	req, err := http.NewRequest(http.MethodGet, APIURL, nil)

	req.Header.Add("authorization", "token "+token)
	req.Header.Add("accept", "application/json")
	req.Header.Add("content-type", "application/json")
     
	if err != nil {
		panic(err)
	}
	fmt.Println(req)
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
	fmt.Println(string(body))
}
