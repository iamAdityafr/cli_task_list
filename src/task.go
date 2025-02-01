package main

import (
	"errors"
	"fmt"
	"os"
	"time"

	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
)

type Task struct {
	Title       string
	Finished    bool
	CreatedAt   time.Time
	CompletedAt *time.Time
}

type Tasks []Task

func (todos *Tasks) Add(title string) {
	todo := Task{
		Title:       title,
		Finished:    false,
		CompletedAt: nil,
		CreatedAt:   time.Now(),
	}

	*todos = append(*todos, todo)
}

func (todos *Tasks) valInd(ind int) error {
	if ind < 0 || ind >= len(*todos) {
		err := errors.New("invalid index")
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

	isCompleted := t[ind].Finished

	if !isCompleted {
		completionTime := time.Now()
		t[ind].CompletedAt = &completionTime
	}

	t[ind].Finished = !isCompleted
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

type taskItem struct {
	index int
	task  Task
}

func (i taskItem) Title() string {
	status := "❌"
	if i.task.Finished {
		status = "✅"
	}
	return fmt.Sprintf("%d. %s %s", i.index, i.task.Title, status)
}

func (i taskItem) Description() string {
	created := i.task.CreatedAt.Format(time.RFC1123)
	completed := ""
	if i.task.CompletedAt != nil {
		completed = i.task.CompletedAt.Format(time.RFC1123)
	}
	return fmt.Sprintf("Created: %s | Completed: %s", created, completed)
}

func (i taskItem) FilterValue() string {
	return i.task.Title

}

type taskListModel struct {
	list list.Model
}

func (m taskListModel) Init() tea.Cmd {
	return nil
}

func (m taskListModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "q", "ctrl+c":
			return m, tea.Quit
		}
	}

	var cmd tea.Cmd
	m.list, cmd = m.list.Update(msg)
	return m, cmd
}

func (m taskListModel) View() string {
	return m.list.View()
}

func (todos *Tasks) print() {
	var items []list.Item
	for i, task := range *todos {
		items = append(items, taskItem{index: i, task: task})
	}

	l := list.New(items, list.NewDefaultDelegate(), 80, 20)
	l.Title = "Tasks"

	m := taskListModel{list: l}
	p := tea.NewProgram(m)

	if err := p.Start(); err != nil {
		fmt.Println("Error running program:", err)
		os.Exit(1)
	}
}
