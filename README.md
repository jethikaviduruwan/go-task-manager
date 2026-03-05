# go-task-manager

A fast, minimal CLI task manager written in Go. Stores tasks locally in `~/.taskman/tasks.json`.

## Features
- Add tasks with priority levels (low / medium / high)
- Tag tasks for easy filtering
- Mark tasks as done, delete individually, or bulk-clear completed
- Persistent JSON storage

## Installation

```bash
go install github.com/jethikaviduruwan/go-task-manager@latest
```

Or build from source:

```bash
git clone https://github.com/jethikaviduruwan/go-task-manager
cd go-task-manager
go build -o taskman .
```

## Usage

```bash
# Add a task
taskman add "Write unit tests" --priority high --tags "dev,testing"

# List all tasks
taskman list

# Mark done
taskman done 1

# Delete a task
taskman delete 2

# Clear all completed tasks
taskman clear
```

## Storage

Tasks are stored in `~/.taskman/tasks.json` — portable and easy to back up.
// v1-1
// v4-0
// v7-1
// v9-2
// v12-0
// v15-1
