package main

import (
	"fmt"
	"strings"
	"time"

	"github.com/spf13/cobra"
)

var (
	flagPriority string
	flagTags     string
)

var addCmd = &cobra.Command{
	Use:   "add [title]",
	Short: "Add a new task",
	Args:  cobra.MinimumNArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		store, err := loadStore()
		if err != nil {
			return err
		}
		title := strings.Join(args, " ")
		p := Priority(flagPriority)
		var tags []string
		if flagTags != "" {
			for _, t := range strings.Split(flagTags, ",") {
				tags = append(tags, strings.TrimSpace(t))
			}
		}
		task := store.AddTask(title, p, tags)
		if err := store.save(); err != nil {
			return err
		}
		fmt.Printf("✓ Added task #%d: %s [%s]\n", task.ID, task.Title, task.Priority)
		return nil
	},
}

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List all tasks",
	RunE: func(cmd *cobra.Command, args []string) error {
		store, err := loadStore()
		if err != nil {
			return err
		}
		if len(store.Tasks) == 0 {
			fmt.Println("No tasks yet. Add one with: taskman add <title>")
			return nil
		}
		fmt.Printf("%-4s %-6s %-8s %-40s %s\n", "ID", "DONE", "PRIORITY", "TITLE", "TAGS")
		fmt.Println(strings.Repeat("─", 72))
		for _, t := range store.Tasks {
			done := "[ ]"
			if t.Done {
				done = "[✓]"
			}
			tags := strings.Join(t.Tags, ", ")
			fmt.Printf("%-4d %-6s %-8s %-40s %s\n", t.ID, done, t.Priority, t.Title, tags)
		}
		pending := store.Pending()
		fmt.Printf("\n%d pending, %d total\n", len(pending), len(store.Tasks))
		return nil
	},
}

var doneCmd = &cobra.Command{
	Use:   "done [id]",
	Short: "Mark a task as done",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		store, err := loadStore()
		if err != nil {
			return err
		}
		var id int
		fmt.Sscan(args[0], &id)
		task := store.FindByID(id)
		if task == nil {
			return fmt.Errorf("task #%d not found", id)
		}
		now := time.Now()
		task.Done = true
		task.DoneAt = &now
		if err := store.save(); err != nil {
			return err
		}
		fmt.Printf("✓ Task #%d marked as done: %s\n", task.ID, task.Title)
		return nil
	},
}

var deleteCmd = &cobra.Command{
	Use:   "delete [id]",
	Short: "Delete a task",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		store, err := loadStore()
		if err != nil {
			return err
		}
		var id int
		fmt.Sscan(args[0], &id)
		if !store.DeleteByID(id) {
			return fmt.Errorf("task #%d not found", id)
		}
		if err := store.save(); err != nil {
			return err
		}
		fmt.Printf("✓ Task #%d deleted\n", id)
		return nil
	},
}

var clearCmd = &cobra.Command{
	Use:   "clear",
	Short: "Remove all completed tasks",
	RunE: func(cmd *cobra.Command, args []string) error {
		store, err := loadStore()
		if err != nil {
			return err
		}
		n := store.ClearDone()
		if err := store.save(); err != nil {
			return err
		}
		fmt.Printf("✓ Cleared %d completed task(s)\n", n)
		return nil
	},
}

func init() {
	addCmd.Flags().StringVarP(&flagPriority, "priority", "p", "medium", "Task priority (low|medium|high)")
	addCmd.Flags().StringVarP(&flagTags, "tags", "t", "", "Comma-separated tags")
}
// v1-0
// v3-1
// v7-0
// v9-1
// v11-1
// v15-0
