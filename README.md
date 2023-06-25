```
# Fiber Client-Side Caching CRUD API

This is a simple API built with Go and Fiber framework to demonstrate client-side caching. The API provides endpoints for managing a task collection.

## Features

- `GET /api/tasks` endpoint retrieves a list of tasks. The response is cached on the client-side for a specific duration.
- `GET /api/tasks/:id` endpoint retrieves details of a specific task. The response is cached on the client-side for a specific duration.
- `POST /api/tasks` endpoint adds a new task to the collection.
- `PUT /api/tasks/:id` endpoint updates an existing task.
- `DELETE /api/tasks/:id` endpoint deletes a task from the collection.
```
## Installation

1. Clone the repository:

```bash
git clone <repository-url>
```

2. Navigate to the project directory:

```bash
cd fiber-clientside-cache
```

3. Install the dependencies:

```bash
go mod download
```

4. Start the server:

```bash
go run main.go
```

The server will be available at `http://localhost:3000`.

## Usage

You can use tools like cURL or Postman to interact with the API endpoints. Here are some examples:

- Get all tasks:

```bash
curl http://localhost:3000/api/tasks
```

- Get a specific task (replace `<task-id>` with the actual task ID):

```bash
curl http://localhost:3000/api/tasks/<task-id>
```

- Add a new task:

```bash
curl -X POST -H "Content-Type: application/json" -d '{"title":"New Task","description":"Task description"}' http://localhost:3000/api/tasks
```

- Update an existing task (replace `<task-id>` with the actual task ID):

```bash
curl -X PUT -H "Content-Type: application/json" -d '{"title":"Updated Task"}' http://localhost:3000/api/tasks/<task-id>
```

- Delete a task (replace `<task-id>` with the actual task ID):

```bash
curl -X DELETE http://localhost:3000/api/tasks/<task-id>
```

## Configuration

You can modify the caching duration and other settings by editing the code in `main.go`.

