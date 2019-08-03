package main

import (
    "fmt"
    "io/ioutil"
    "net/http"
    "os"
)

func main() {
    
    searchRepo := os.Args[1:]
    APIURL := "http://gitea-tooling.az.globaldevops.gdpdentsu.net/api/v1/admin/orgs?access_token=b479bc9baa5ae757a458d73b2a610eafce6e4d5a"
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
    fmt.Printf("%v", string(body))
}

