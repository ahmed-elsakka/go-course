package cmd

import (
	"flag"
	"fmt"
	"tasks-management/models"
	"time"
)

func RunCLI(repo models.TaskRepository) {
	addCmd := flag.NewFlagSet("add", flag.ExitOnError)
	title := addCmd.String("title", "", "Task title")
	desc := addCmd.String("desc", "", "Task description")

	listCmd := flag.NewFlagSet("list", flag.ExitOnError)

	if len(flag.Args()) < 1 {
		fmt.Println("expected 'add' or 'list' subcommands")
		return
	}

	switch flag.Arg(0) {
	case "add":
		addCmd.Parse(flag.Args()[1:])
		task := &models.Task{
			Title:       *title,
			Description: *desc,
			Status:      "Pending",
			CreatedAt:   time.Now(),
		}
		repo.Create(task)
		fmt.Println("✅ Task added:", task.Title)

	case "list":
		listCmd.Parse(flag.Args()[1:])
		tasks, _ := repo.GetAll(1)
		for _, t := range tasks {
			fmt.Printf("[%d] %s - %s\n", t.ID, t.Title, t.Status)
		}
	}
}
