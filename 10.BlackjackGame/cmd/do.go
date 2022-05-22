package cmd

/*
import (
	taskrepository "clitaskmanager/taskrepository"
	"fmt"
	"log"
	"strings"

	"github.com/spf13/cobra"
)

// doCmd represents the do command
var doCmd = &cobra.Command{
	Use:   "do",
	Short: "'Does' a task on the task list, removing it from the todo list.",
	Long: `Removes a task from the task list. If no task is found with the same name,
	will return a message saying that the task could not be found.`,
	Run: func(cmd *cobra.Command, args []string) {
		taskName := strings.Join(args, " ")
		fmt.Printf("Removing the following task from task list: %v\n", taskName)
		err := taskrepository.Delete(taskName)
		if err != nil {
			log.Fatal(err.Error())
		}
	},
}

func init() {
	rootCmd.AddCommand(doCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// doCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// doCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
*/
