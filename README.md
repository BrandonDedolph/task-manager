# Task Manager

A terminal-based task manager built with Go and [Bubble Tea](https://github.com/charmbracelet/bubbletea).

![Go](https://img.shields.io/badge/Go-1.26-00ADD8?logo=go&logoColor=white)
![License](https://img.shields.io/badge/license-MIT-blue)

## Features

- **Create** tasks with a name, description, status, and priority
- **Edit** existing tasks inline
- **Delete** tasks with a confirmation prompt
- **List** all tasks with status indicators
- **Persist** tasks to a local JSON file

## Architecture

The project follows clean architecture principles with clearly separated layers:

```
cmd/
└── main.go          # Entry point — wires dependencies together

internal/
├── task/            # Domain model and repository interface
├── datastore/       # File-based JSON persistence (implements Repository)
├── service/         # Business logic (CRUD operations)
└── tui/             # Terminal UI (list, form, confirm views)
```

**Data flow:** `TUI → TaskService → FileRepository → tasks.json`

## Tech Stack

- [Bubble Tea](https://github.com/charmbracelet/bubbletea) — TUI framework
- [Bubbles](https://github.com/charmbracelet/bubbles) — UI components (text input)
- [google/uuid](https://github.com/google/uuid) — Task ID generation

## Getting Started

### Prerequisites

- Go 1.26+

### Install

```sh
go install github.com/BrandonDedolph/task-manager@latest
```

### Build from source

```sh
git clone https://github.com/BrandonDedolph/task-manager.git
cd task-manager
go build -o task-manager ./cmd/...
```

### Run

```sh
go run ./cmd/
```

## Usage

| Key              | Action                  |
|------------------|-------------------------|
| `j` / `↓`       | Move cursor down        |
| `k` / `↑`       | Move cursor up          |
| `a`              | Add new task            |
| `e`              | Edit selected task      |
| `d`              | Delete selected task    |
| `Space` / `Enter`| Complete selected task  |
| `Enter`          | Confirm / submit form   |
| `Esc`            | Cancel / go back        |
| `y`              | Confirm deletion        |
| `n`              | Cancel deletion         |
| `q` / `Ctrl+C`  | Quit                    |

## Task Model

Each task has the following fields:

| Field         | Type     | Description                              |
|---------------|----------|------------------------------------------|
| `id`          | string   | Auto-generated UUID                      |
| `name`        | string   | Task name                                |
| `description` | string   | Optional description                     |
| `status`      | string   | `todo`, `inprogress`, or `complete`      |
| `priority`    | int      | Priority level                           |
