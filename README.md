# Task Management CRUD Application

A full-stack task management application built with Vue.js frontend and Go backend.

## Features

- Create, read, update, and delete tasks
- Task status management (pending, in-progress, completed)
- Responsive UI design

## Project Structure

```
testgolangapp/
├── server/                  # Go backend
│   ├── cmd/api/             # API entry point
│   └── internal/            # Internal packages
│       ├── database/        # Database operations
│       ├── handlers/        # HTTP handlers
│       └── models/          # Data models
├── src/                     # Vue.js frontend
│   ├── components/          # Vue components
│   ├── services/            # API services
│   ├── views/               # Vue views
│   ├── App.vue              # Root component
│   └── main.js              # Entry point
└── tasks.db                 # SQLite database
```

## Prerequisites

- Node.js and npm/yarn
- Go 1.24 or higher
- SQLite

## Startup Procedure

### Frontend Setup

1. Install dependencies:
   ```
   yarn install
   ```

2. Start development server:
   ```
   yarn serve
   ```
   The frontend will be available at http://localhost:8080

### Backend Setup

1. Navigate to the server directory:
   ```
   cd server/cmd/api
   ```

2. Run the Go server:
   ```
   go run main.go
   ```
   The backend API will be available at http://localhost:8000

## API Endpoints

- `GET /api/tasks` - Get all tasks
- `GET /api/tasks/{id}` - Get a specific task
- `POST /api/tasks` - Create a new task
- `PUT /api/tasks/{id}` - Update a task
- `DELETE /api/tasks/{id}` - Delete a task

## Development

### Linting

```
yarn lint
```

or

```
yarn eslint
```

### Building for Production

```
yarn build
```

## Database

The application uses SQLite for data storage. The database file `tasks.db` is created automatically when the backend server starts.
