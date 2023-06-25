package main

import (
	"encoding/json"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"time"
)

type Task struct {
	ID          uuid.UUID `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type TaskRequest struct {
	Title       string `json:"title"`
	Description string `json:"description"`
}

type CacheItem struct {
	Data      interface{}
	ExpiresAt time.Time
}

var (
	tasks     = make(map[uuid.UUID]Task)
	taskCount = 0
)

func main() {

	app := fiber.New()
	group := app.Group("/api")
	group.Post("/tasks", CreateTask)
	group.Get("/tasks/:id", GetTaskByID)
	group.Get("/tasks", ListTasks)
	group.Delete("/tasks/:id", DeleteTaskByID)
	group.Put("/tasks/:id", UpdateTaskByID)

	err := app.Listen(":3000")
	if err != nil {
		panic(err)
	}
}

func ListTasks(c *fiber.Ctx) error {
	cacheKey := "tasks_cache_key"
	if c.Get(cacheKey) != "" {
		return c.JSON(c.Get(cacheKey))
	}
	tasksList := getTasksList()
	cacheItem := &CacheItem{
		Data:      tasksList,
		ExpiresAt: time.Now().Add(5 * time.Minute),
	}
	out, err := json.Marshal(cacheItem)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).
			SendString("An error occurred when marshalling the cache item")
	}

	c.Set(cacheKey, string(out))
	return c.JSON(fiber.Map{"data": tasksList, "count": taskCount})
}

func CreateTask(c *fiber.Ctx) error {
	var payload TaskRequest
	if err := c.BodyParser(&payload); err != nil {
		return c.Status(fiber.StatusBadRequest).
			SendString("An error occurred when parsing the request body")
	}
	task := Task{
		ID:          uuid.New(),
		Title:       payload.Title,
		Description: payload.Description,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}
	taskCount++
	tasks[task.ID] = task
	c.Set("tasks_cache_key", "")
	return c.JSON(task)
}

func GetTaskByID(c *fiber.Ctx) error {
	id := c.Params("id")
	if c.Get(id) != "" {
		return c.JSON(c.Get(id))
	}
	s, err := uuid.Parse(id)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).
			SendString("An error occurred when parsing the uuid")
	}
	if task, ok := tasks[s]; ok {
		cacheItem := &CacheItem{
			Data:      task,
			ExpiresAt: time.Now().Add(5 * time.Minute),
		}
		out, err := json.Marshal(cacheItem)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).
				SendString("An error occurred when marshalling the cache item")
		}
		c.Set(id, string(out))
		return c.JSON(task)
	}
	return c.Status(fiber.StatusNotFound).SendString("Task not found")
}

func DeleteTaskByID(c *fiber.Ctx) error {
	id := c.Params("id")
	s, err := uuid.Parse(id)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).
			SendString("An error occurred when parsing the uuid")
	}
	if _, ok := tasks[s]; ok {
		delete(tasks, s)
		taskCount--
		c.Set(id, "")
		return c.SendString("Task deleted successfully")
	}
	return c.Status(fiber.StatusNotFound).SendString("Task not found")
}

func UpdateTaskByID(c *fiber.Ctx) error {
	id := c.Params("id")
	s, err := uuid.Parse(id)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).
			SendString("An error occurred when parsing the uuid")
	}
	var updatedTask Task
	var payload TaskRequest
	if err := c.BodyParser(&payload); err != nil {
		return c.Status(fiber.StatusBadRequest).
			SendString("An error occurred when parsing the request body")
	}

	if _, ok := tasks[s]; ok {
		updatedTask = Task{
			ID:          tasks[s].ID,
			Title:       payload.Title,
			Description: payload.Description,
			CreatedAt:   tasks[s].CreatedAt,
			UpdatedAt:   time.Now(),
		}
		tasks[s] = updatedTask
		c.Set(id, "")
		return c.JSON(updatedTask)
	}
	return c.Status(fiber.StatusNotFound).SendString("Task not found")
}

func getTasksList() []Task {
	taskList := make([]Task, 0, len(tasks))
	for _, task := range tasks {
		taskList = append(taskList, task)
	}
	return taskList
}
