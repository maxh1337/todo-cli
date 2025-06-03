# 📝 todo-cli

A simple and colorful command-line ToDo manager written in Go, built with [Cobra](https://github.com/spf13/cobra).  
Stylized output using custom color utilities for a better terminal experience.

## 📦 Installation

Make sure you have Go installed (`go version`), then run:

```bash
git clone https://github.com/yourusername/todo-cli.git
cd todo-cli
go build -o todo
```

This will compile a binary named `todo`.

> Optional: move it to your PATH for global access
> `mv todo /usr/local/bin/`

## 🚀 Usage

```bash
todo <command> [arguments...]
```

### ✍️ Create a new task

```bash
todo create "Buy groceries"
```

### 📋 List all tasks

```bash
todo list
```

### ✅ Mark a task as done

```bash
todo done 1
```

### ❌ Delete a task

```bash
todo delete 1
```

## 🎨 Example Output

```bash
$ todo list

[1] Buy groceries         ☐
[2] Finish the report     ✅
[3] Call Alice            ☐
```

Tasks are displayed with colored labels and checkmarks for clarity.

## 🧱 Project Structure

```
.
├── cmd/                  # Command implementations (Cobra)
│   ├── create.go
│   ├── delete.go
│   ├── done.go
│   ├── list.go
│   └── root.go
├── internal/
│   ├── app/              # Application core logic
│   │   └── app.go
│   └── storage/          # Storage layer (JSON file-based)
│       ├── storage.go
│       └── todos.json
├── utils/
│   └── colors.go         # Terminal color helpers
├── main.go               # Entry point
├── go.mod
├── go.sum
├── todo                  # Compiled binary (optional)
└── project-structure.txt # Text description of layout
```

## 💾 Persistence

Tasks are stored in a local JSON file:
`internal/storage/todos.json`
This allows you to keep your to-do list between sessions.
