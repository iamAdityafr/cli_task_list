# CLI Terminal Task List

A Command Line Interface (CLI) task management application built with **Go** and the **Bubble Tea** framework. Organize and manage your tasks directly from the terminal with a delightful TUI  experience.

---

## Features

- **Add Tasks**: Add new tasks.
- **List Tasks**: View all your tasks in a clean, interactive list.
- **Mark Tasks**: Easily mark tasks.
- **Delete Tasks**: Remove tasks.
- **Interactive TUI**: Enjoy a smooth and responsive terminal interface powered by Bubble Tea.
- **Persistent Storage**: Tasks are saved to a local file and loaded automatically when the app starts.
- **Filter**: Searching for the tasks.

---

## Installation

1. **Clone the repository**:
   ```bash
   git clone https://github.com/iamAdityafr/cli_task_list
   cd cli_task_list

2. Build it:

```bash
go build -o tasklist.exe
```
3. Run it:

```bash
tasklist
```
## Screenshots

![App Screenshot1](https://github.com/iamAdityafr/cli_task_list/blob/main/src/img/ig1.png)

---

![App Screenshot2](https://github.com/iamAdityafr/cli_task_list/blob/main/src/img/ig2.png)

## Usage
Below are the available commands and their usage:

```bash
1. Add a Task
## To add a new task, use the -add flag followed by the task title:


tasklist -add "Finish Project"


2. Edit a Task
## To edit an existing task, use the -edit flag followed by the task index and the new title in the format index:new_title:

tasklist -edit 1:Go exercise


3. Delete a Task
## To delete a task, use the -del flag followed by the task index:

tasklist -del 1


4. Toggle Task Completion
## To toggle the completion status of a task, use the -toggle flag followed by the task index:

tasklist -toggle 1


5. List All Tasks
## To list all tasks, use the -list flag:

tasklist -list

```
