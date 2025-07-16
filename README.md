# todo-cli

A simple and efficient command-line todo application written in Go.

---

## Features

- **Add** new tasks to your todo list
- **List** unfinished or all tasks
- **Mark** tasks as completed or undo completion
- **Delete** tasks by ID
- Fast, lightweight, and easy to use

---

## Installation

### 1. Build from Source

```sh
git clone https://github.com/yourusername/todo-cli.git
cd todo-cli
go build -o todo-cli
```

### 2. (Optional) Move the Binary to Your PATH

```sh
mv todo-cli /usr/local/bin/
```

---

## Usage

All commands are available via the `todo-cli` executable.

### Add a Task

```sh
todo-cli add "Buy groceries"
```
_Adds a new task with the title "Buy groceries"._

---

### List Tasks

```sh
todo-cli list
```
_Lists all unfinished tasks._

```sh
todo-cli list --all
# or
todo-cli list -a
```
_Lists all tasks, including completed ones._

---

### Complete a Task

```sh
todo-cli complete 3
```
_Marks the task with ID 3 as completed._

```sh
todo-cli complete 3 --reverse
# or
todo-cli complete 3 -r
```
_Marks the task with ID 3 as unfinished (undo completion)._

---

### Delete a Task

```sh
todo-cli delete 5
```
_Deletes the task with ID 5 from your todo list._

---

## Examples

```sh
todo-cli add "Read a book"
todo-cli list
todo-cli complete 1
todo-cli list
todo-cli complete 1 --reverse
todo-cli delete 1
```

---

## Help

For help on any command, use the `--help` flag:

```sh
todo-cli --help
todo-cli add --help
todo-cli list --help
```

---

## Configuration

- Uses an SQLite database by default.
- Set the database location with the `DB_URL` environment variable.

---

## Credits

Built with [Cobra](https://github.com/spf13/cobra) and Go.

AI Generated README using [VaultAI](https://app.vaultai.eu) (GPT-4.1).

---

## Author

* Pierre Teodoresco

---

*Happy productivity!*
