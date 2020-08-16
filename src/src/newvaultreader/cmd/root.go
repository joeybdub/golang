package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var (
	resource string
)

var rootCmd = &cobra.Command{
	Use:   "resoure",
	Short: "azure resource required",
	Long: `Enter the resource type that is required 
          The is a full resoure on Azure
                Complete documentation is available at https://docs.microsoft.com/en-us/rest/api/resources/`,
	Run: func(cmd *cobra.Command, args []string) {
		// Do Stuff Here
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
