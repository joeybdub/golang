package azuresub

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func AzureSub() string {
	text := bufio.NewReader(os.Stdin)
	fmt.Print("Enter subscription: ")
	Subscription, err := text.ReadString('\n')
    Subscription = strings.TrimSuffix(Subscription, "\n")
	check(err)

	return Subscription

}
