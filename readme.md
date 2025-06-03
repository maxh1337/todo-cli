# 📝 todo-cli – Two CLI ToDo Apps in Go

This repository contains two separate implementations of a terminal-based ToDo application written in Go:

1. **todo-cobra** – feature-rich version using [Cobra](https://github.com/spf13/cobra) and color output
2. **todo-basic** – minimal implementation using native `os`, `flag`, and `fmt` packages only

Ideal for learning how to build CLI tools in Go — from scratch or with a powerful CLI framework.

---

## 📁 Project Structure

```

todo-cli/
├── todo-cobra/     # Implementation with Cobra + colored output
│   ├── cmd/
│   ├── internal/
│   ├── utils/
│   └── main.go
│   └── README.md
│
├── todo-basic/     # Basic implementation using only standard library
│   ├── storage/
│   └── main.go
│   └── README.md
│
└── README.md       # This file (overview of both versions)

```

---

## 🚀 Version 1 – todo-cobra

**Features:**

- Interactive CLI powered by Cobra
- Colored output with ✅ / ☐ indicators
- Subcommands: `create`, `list`, `done`, `delete`
- Task persistence via local JSON storage

### 📦 Run

```bash
cd todo-cobra
go build -o todo
./todo create "Read a book"
./todo list
```

More info: [`todo-cobra/README.md`](./todo-cobra/README.md)

---

## ⚙️ Version 2 – todo-basic

**Features:**

- No external dependencies
- Simple interface via `flag` or positional arguments
- JSON-based storage
- Ideal for understanding low-level CLI processing

### 📦 Run

```bash
cd todo-basic
go build -o todo
./todo -create "Buy milk"
./todo -list
```

Available flags:

- `-create "task title"`
- `-list`
- `-done 2`
- `-delete 1`

More info: [`todo-basic/README.md`](./todo-basic/README.md)

---

## 🤔 Why Two Versions?

| Feature        | todo-cobra      | todo-basic     |
| -------------- | --------------- | -------------- |
| CLI Framework  | ✅ Cobra        | ❌ stdlib only |
| Colors         | ✅ With `color` | ❌ Plain text  |
| Beginner level | ⭐⭐            | ⭐             |
| Extensibility  | ⭐⭐⭐⭐        | ⭐             |
