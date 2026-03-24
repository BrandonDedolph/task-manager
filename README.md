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
internal/
├── task/          # Domain model and repository interface
├── service/       # Business logic (CRUD operations)
├── datastore/     # File-based JSON persistence
└── tui/           # Terminal UI (list, form, confirm views)
```

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
./task-manager
```

## Usage

| Key          | Action              |
|--------------|---------------------|
| `j` / `↓`   | Move cursor down    |
| `k` / `↑`   | Move cursor up      |
| `Enter`      | Confirm / submit    |
| `Esc`        | Cancel / go back    |
| `y`          | Confirm deletion    |
| `n`          | Cancel deletion     |
| `q`          | Quit                |

## Task Model

Each task has the following fields:

| Field         | Type     | Description                              |
|---------------|----------|------------------------------------------|
| `id`          | string   | Auto-generated UUID                      |
| `name`        | string   | Task name                                |
| `description` | string   | Optional description                     |
| `status`      | string   | `todo`, `inprogress`, or `complete`      |
| `priority`    | int      | Priority level                           |
