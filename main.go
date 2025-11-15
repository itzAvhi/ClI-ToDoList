package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/spf13/cobra"
)

// Task structure
type task struct {
	name   string
	status string
}

var todolist []task

// Root command
var rootCmd = &cobra.Command{
	Use:   "todolist",
	Short: "A simple CLI To-Do-List",
	Long:  `A simple CLI To-Do-List built with cobra for further CLI commands`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Welcome to todolist! Type 'todolist --help' for available commands.")
	},
}

// Adding a command
var addCmd = &cobra.Command{
	Use:   "add [task]",
	Short: "Add a new task",
	Long:  `Add a new task to your todo list`,
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		taskName := strings.Join(args, " ")
		todolist = append(todolist, task{
			name:   taskName,
			status: "[ ]",
		})
		fmt.Printf("✓ Added task: %s\n", taskName)
		fmt.Printf("Total tasks: %d\n", len(todolist))
	},
}

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "View all tasks",
	Long:  `Display all tasks in your todo list`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(todolist) == 0 {
			fmt.Println("No tasks yet! Add one with: todolist add [task]")
			return
		}

		fmt.Println("\nYour Tasks:")
		fmt.Println("===========")
		for i, t := range todolist {
			fmt.Printf("%d. %s %s\n", i+1, t.status, t.name)
		}
		fmt.Println()
	},
}

var doneCmd = &cobra.Command{
	Use:   "done [task number]",
	Short: "Mark a task as done",
	Long:  `Mark a task as completed by its number`,
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		var taskNum int
		_, err := fmt.Sscanf(args[0], "%d", &taskNum)

		if err != nil {
			fmt.Println("Error: Please provide a valid task number")
			return
		}

		if taskNum < 1 || taskNum > len(todolist) {
			fmt.Printf("Error: Task number must be between 1 and %d\n", len(todolist))
			return
		}

		todolist[taskNum-1].status = "[✓]"
		fmt.Printf("✓ Marked task %d as done: %s\n", taskNum, todolist[taskNum-1].name)
	},
}

var deleteCmd = &cobra.Command{
	Use:   "delete [task number]",
	Short: "Delete a task",
	Long:  `Delete a task from your todo list by its number`,
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		var taskNum int
		_, err := fmt.Sscanf(args[0], "%d", &taskNum)

		if err != nil {
			fmt.Println("Error: Please provide a valid task number")
			return
		}

		if taskNum < 1 || taskNum > len(todolist) {
			fmt.Printf("Error: Task number must be between 1 and %d\n", len(todolist))
			return
		}

		deletedTask := todolist[taskNum-1].name

		todolist = append(todolist[:taskNum-1], todolist[taskNum:]...)
		fmt.Printf("✓ Deleted task: %s\n", deletedTask)
		fmt.Printf("Remaining tasks: %d\n", len(todolist))
	},
}

func init() {
	// Register all commands
	rootCmd.AddCommand(addCmd)
	rootCmd.AddCommand(listCmd)
	rootCmd.AddCommand(doneCmd)
	rootCmd.AddCommand(deleteCmd)
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
