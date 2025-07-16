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

## Database Setup

This project uses SQLite with [sqlc](https://sqlc.dev/) for type-safe SQL code generation and [golang-migrate](https://github.com/golang-migrate/migrate) for database migrations.

### Prerequisites

Install the required tool:

```sh
# Install golang-migrate
go install -tags 'sqlite3' github.com/golang-migrate/migrate/v4/cmd/migrate@latest
```

### Setup Database

```sh
# Create and migrate the database
migrate -path migrations -database "sqlite3://todo.db" up
```

### Database Configuration

- **Default**: Uses `todo.db` SQLite file in the current directory
- **Custom location**: Set the `DB_URL` environment variable

```sh
export DB_URL="sqlite3://path/to/your/database.db"
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

## Credits

Built with [Cobra](https://github.com/spf13/cobra) and Go.

README co-written with AI using [VaultAI](https://app.vaultai.eu) (GPT-4.1) and [Cursor](https://www.cursor.com/) (claude-4-sonnet).

---

## Author

- Pierre Teodoresco

---

_Happy productivity!_
