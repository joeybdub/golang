package main

import (
	"fmt"
	"github.com/hashicorp/vault/api"
	"os"
)

type authConfig struct {
    token     string `yaml:"token"`
    vault_addr string `yaml:"vault_addr"`
    path     string `yaml:"path"`
}


func main() {
	config := &api.Config{
		Address: vault_addr,
	}
	client, err := api.NewClient(config)
	if err != nil {
		fmt.Println(err)
		return
	}
	client.SetToken(token)
	c := client.Logical()
	secret, err := c.Read("t/")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(secret.Data["hello"])
}
