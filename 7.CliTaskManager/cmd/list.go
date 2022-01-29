package cmd

import (
	taskrepository "clitaskmanager/taskrepository"
	"fmt"
	"log"

	"github.com/spf13/cobra"
)

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "Lists all tasks currently on the task list.",
	Long:  `Lists all tasks currently on the task list. Does not accept any arguments.`,
	Run: func(cmd *cobra.Command, args []string) {
		tasks, err := taskrepository.Read()
		if err != nil {
			log.Fatal(err.Error())
		}

		for _, value := range tasks {
			fmt.Println(value)
		}
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
}
