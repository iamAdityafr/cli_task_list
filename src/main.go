package main

func main() {
	todos := Tasks{}
	storage := NewStorage[Tasks]("todos.json")
	storage.Load(&todos)
	cmd := NewFlags()
	cmd.Execute(&todos)
	todos.print()
	storage.Save(todos)

}
