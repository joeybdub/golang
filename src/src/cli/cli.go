package main

import (
	"fmt"
	"log"
	"net/url"
	//"net/http"
	"os"
	//"io/ioutil"
)

func main() {
   	//cmdArgs := os.Args[1:]
	searchRepo := os.Args[1]
	giteaToken := os.Args[2]
	giteaInstance := os.Args[4]
	apiPath :=  os.Args[3]
	//path,err := "api/v1/repos/search?q='$searchRepo'&access_token=b479bc9baa5ae757a458d73b2a610eafce6e4d5a"
	
	urlConstruct := "api/v1/"+ apiPath +"?q"+ searchRepo +"&access_token="+ giteaToken
	//fmt.Print(urlConstruct)
	//path, err := url.Parse("api/v1/repos/search?q=Joe&access_token=b479bc9baa5ae757a458d73b2a610eafce6e4d5a")
  
    path, err := url.Parse(urlConstruct)	
	  if err != nil {
		log.Fatal(err)
	}
	url := url.URL{
		Scheme: "https",
		Host:   "gitea-tooling.az."+giteaInstance+".gdpdentsu.net",
	}
	//url.String() newUrl := url.ResolveReference(path)
	EndURL := url.ResolveReference(path) 
	fmt.Println(EndURL)
 //   return url.ResolveReference(path)
}
