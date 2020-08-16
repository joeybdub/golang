package main

import (
	"fmt"

	"github.com/plurasight/webservice/models"
)

func main() {

	u := models.User{
		ID:        2,
		FirstName: "Joe",
		LastName:  "Beadle",
	}

	fmt.Println(u)
}
