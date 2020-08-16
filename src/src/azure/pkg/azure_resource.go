package cmd
import (
	"fmt"
	"io/ioutil"
	"log"
	"path"
        "github.com/spf13/viper"
	"github.com/spf13/cobra"
)

var (
    acr string
    storage string
) 
    

var resourceCmd = &cobra.Command{
	Use:   "resource",
	Short: "Run an resource on azure",
	Long:  `This will apply premissions to the azure ACR's. `,
	PreRun: func(cmd *cobra.Command, args []string) {
		common.LogColour("heading", "INFO: Running a plan on your payload")
	},
	Run: func(cmd *cobra.Command, args []string) {
		fflags := cmd.Flags()                // fflags is a *flag.FlagSet
		if fflags.Changed("type") == false { // check if the flag "path" is set
			fmt.Println("You need to specify a resource type") // If not, we'll let the user know
			return                                    // and return
		}
		if err != nil {
			log.Fatal(err)
		}

		azure.acr_update(type, remove, update)
	},
}


func init() {
	rootCmd.AddCommand(resourceCmd)
	resourceCmd.PersistentFlags().StringVar(&type, "type", "", "The type of resource you are looking to update in azure")
	resourceCmd.Flags().BoolVar(&remnove, "remove", false, "Whether to remove access azure resource")
	resourceCmd.Flags().BoolVar(&update, "update", true, "Whether to to update terraform modules")
        resourceCmd.Flags().BoolVar(&update, "update", true, "List resource details") 

}


