package main

import (
	azuresub "azure/pkg"
	"fmt"
	"os"

	"github.com/hashicorp/vault/api"
)

var token = os.Getenv("TOKEN")
var vault_addr = os.Getenv("VAULT_ADDR")

func main() {
	sub123 := azuresub.AzureSub()
	config := &api.Config{
		Address: vault_addr,
	}
	client, err := api.NewClient(config)
	if err != nil {
		fmt.Println(err)
		return
	}
	client.SetToken(token)
	secret, err := client.Logical().Read("secret/accounts/" + sub123 + "/azure")
	if err != nil {
		fmt.Println(err)
		return
	}
	client_id, _ := secret.Data["azure_client_id"]
	fmt.Println(client_id)
	client_secret, _ := secret.Data["azure_client_secret"]
	fmt.Println(client_secret)

}
