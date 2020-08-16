package main

import (
	"fmt"
	"os"

	"github.com/hashicorp/vault/api"
)

var token = os.Getenv("TOKEN")
var vault_addr = os.Getenv("VAULT_ADDR")

//var sub123 azure_creds.Subscription

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
	secret, err := client.Logical().Read("secret/accounts/devopschina/azure")
	if err != nil {
		fmt.Println(err)
		return
	}
	avalue, _ := secret.Data["client_id"]
	fmt.Println(avalue)
}
