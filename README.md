# Task Management CRUD Application

A simple task management application built with Vue.js frontend and Go backend, using SQLite for data storage.

## Features

- Create, Read, Update, and Delete tasks
- Task status management (Pending, In Progress, Completed)
- RESTful API with Go
- Responsive UI with Vue.js

## Prerequisites

- Go 1.18 or higher
- Node.js and npm/yarn
- SQLite

## Installation

1. Clone the repository:
```bash
git clone https://github.com/KtokKawu/testgolangapp.git
cd testgolangapp
```

2. Install frontend dependencies:
```bash
yarn install
```

3. Install Go dependencies:
```bash
go mod tidy
```

## Running the Application

### Start the Backend Server

```bash
go run server/cmd/api/main.go
```

The backend server will start on http://localhost:8000

### Start the Frontend Development Server

```bash
yarn serve
```

The frontend development server will start on http://localhost:8080

## API Endpoints

| Method | Endpoint | Description |
|--------|----------|-------------|
| GET    | /api/tasks | Get all tasks |
| GET    | /api/tasks/:id | Get a specific task by ID |
| POST   | /api/tasks | Create a new task |
| PUT    | /api/tasks/:id | Update an existing task |
| DELETE | /api/tasks/:id | Delete a task |

## Project Structure

```
testgolangapp/
├── public/                  # Static assets
├── server/                  # Go backend
│   ├── cmd/                 # Entry points
│   │   └── api/             # API server
│   └── internal/            # Internal packages
│       ├── database/        # Database operations
│       ├── handlers/        # HTTP handlers
│       └── models/          # Data models
├── src/                     # Vue.js frontend
│   ├── assets/              # Static assets
│   ├── components/          # Vue components
│   ├── router/              # Vue Router configuration
│   ├── services/            # API services
│   ├── views/               # Vue views
│   ├── App.vue              # Root component
│   └── main.js              # Entry point
└── tasks.db                 # SQLite database
```

## Development

### Linting

```bash
yarn lint
```

For customizing the configuration, see [Vue CLI Configuration Reference](https://cli.vuejs.org/config/).
