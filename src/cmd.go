package main

import (
	"flag"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Flags struct {
	Add    string
	Del    int
	Edit   string
	Toggle int
	List   bool
}

func NewFlags() *Flags {
	cmdflg := Flags{}
	flag.StringVar(&cmdflg.Add, "add", "", "Add a new task with title")
	flag.BoolVar(&cmdflg.List, "list", false, "List all tasks")
	flag.IntVar(&cmdflg.Toggle, "toggle", -1, "Toggle task by index true/false")
	flag.IntVar(&cmdflg.Del, "del", -1, "Delete task by index")
	flag.StringVar(&cmdflg.Edit, "edit", "", "Edit a task by index & set a new title. id:new_title")

	flag.Parse()

	return &cmdflg
}

func (cmdflg *Flags) Execute(todos *Tasks) {
	switch {
	case cmdflg.Add != "":
		todos.Add(cmdflg.Add)
	case cmdflg.List:
		todos.print()
	case cmdflg.Edit != "":
		split := strings.SplitN(cmdflg.Edit, ":", 2)
		if len(split) != 2 {
			fmt.Println("Invalid format for edit. Use index:new_title")
			os.Exit(1)
		}
		index, err := strconv.Atoi(split[0])
		if err != nil {
			fmt.Println("Invalid index for edit.")
			os.Exit(1)

		}
		todos.edit(index, split[1])
	case cmdflg.Toggle != -1:
		todos.toggle(cmdflg.Toggle)

	case cmdflg.Del != -1:
		todos.Delete(cmdflg.Del)

	default:
		fmt.Println("Invalid cmd")
	}
}
