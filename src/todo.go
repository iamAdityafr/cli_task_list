package main

import (
	"errors"
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/aquasecurity/table"
)

type Task struct {
	Title       string
	Completed   bool
	CreatedAt   time.Time
	CompletedAt *time.Time
}

type Tasks []Task

func (todos *Tasks) Add(title string) {
	todo := Task{
		Title:       title,
		Completed:   false,
		CompletedAt: nil,
		CreatedAt:   time.Now(),
	}

	*todos = append(*todos, todo)
}

func (todos *Tasks) valInd(ind int) error {
	if ind < 0 || ind >= len(*todos) {
		err := errors.New("invalid ind")
		fmt.Println(err)
		return err
	}

	return nil
}

func (todos *Tasks) Delete(ind int) error {
	t := *todos

	if err := t.valInd(ind); err != nil {
		return err
	}

	*todos = append(t[:ind], t[ind+1:]...)
	return nil

}

func (todos *Tasks) toggle(ind int) error {
	t := *todos
	if err := t.valInd(ind); err != nil {
		return err
	}

	isCompleted := t[ind].Completed

	if !isCompleted {
		completionTime := time.Now()
		t[ind].CompletedAt = &completionTime
	}

	t[ind].Completed = !isCompleted
	return nil
}

func (todos *Tasks) edit(ind int, title string) error {
	t := *todos
	if err := t.valInd(ind); err != nil {
		return err
	}

	t[ind].Title = title
	return nil
}

func (todos *Tasks) print() {
	table := table.New(os.Stdout)
	table.SetRowLines(false)
	table.SetHeaders("Sr No.", "Task", "isCompleted", "Created", "Completed At")

	for ind, t := range *todos {
		completed := "❌"
		completedAt := ""

		if t.Completed {
			completed = "✅"
			if t.CompletedAt != nil {
				completedAt = t.CompletedAt.Format(time.RFC1123)

			}

		}

		table.AddRow(strconv.Itoa(ind), t.Title, completed, t.CreatedAt.Format(time.RFC1123), completedAt)

	}

	table.Render()
}
