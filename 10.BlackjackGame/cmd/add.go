package cmd

/*
import (
	taskrepository "clitaskmanager/taskrepository"
	"fmt"
	"log"
	"strings"

	"github.com/spf13/cobra"
)

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Adds a task to the task list.",
	Long:  `Adds a task to the task list. The task list is maintained in a local instance of FireboltDB. Accepts a set of strings as an argument.`,
	Run: func(cmd *cobra.Command, args []string) {
		taskName := strings.Join(args, " ")
		fmt.Printf("Adding the following task to the task list: %v\n", taskName)
		err := taskrepository.Update(taskName)
		if err != nil {
			log.Fatal(err.Error())
		}
	},
}

func init() {
	rootCmd.AddCommand(addCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// addCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// addCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
*/
