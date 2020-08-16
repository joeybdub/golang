package buildurl

import (
	//	"fmt"

	"log"
	"net/url"
	"os"
)

func FullURL() *url.URL {
	//	RepoName := flag.String("repositoryName", "Joe", "The name of repositrory to migrate")
	//	GiteaInstance := flag.String("GiteaInstance", "globaldevops", "The gitea instanec your want to connect to")

	searchRepo := os.Args[3]
	giteaToken := os.Args[1]
	giteaInstance := os.Args[4]
	apiPath := os.Args[2]
	urlConstruct := "api/v1/" + apiPath + "?q=" + searchRepo + "&access_token=" + giteaToken

	path, err := url.Parse(urlConstruct)
	if err != nil {
		log.Fatal(err)
	}
	url := url.URL{
		Scheme: "https",
		Host:   "gitea-tooling.az." + giteaInstance + ".gdpdentsu.net",
	}
	EndURL := url.ResolveReference(path)
	return EndURL
}
