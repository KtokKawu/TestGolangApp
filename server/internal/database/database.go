package database

import (
	"database/sql"
	"fmt"
	"log"
	"testgolangapp/server/internal/models"

	_ "github.com/mattn/go-sqlite3"
)

var DB *sql.DB

func InitDB() error {
	var err error
	DB, err = sql.Open("sqlite3", "./tasks.db")
	if err != nil {
		return err
	}

	createTableSQL := `
	CREATE TABLE IF NOT EXISTS tasks (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		title TEXT NOT NULL,
		description TEXT,
		status TEXT DEFAULT 'pending'
	);`

	_, err = DB.Exec(createTableSQL)
	if err != nil {
		return err
	}

	return nil
}

func GetAllTasks() ([]models.Task, error) {
	rows, err := DB.Query("SELECT id, title, description, status FROM tasks")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var tasks []models.Task
	for rows.Next() {
		var task models.Task
		err := rows.Scan(&task.ID, &task.Title, &task.Description, &task.Status)
		if err != nil {
			return nil, err
		}
		tasks = append(tasks, task)
	}

	return tasks, nil
}

func GetTask(id int) (models.Task, error) {
	var task models.Task
	err := DB.QueryRow("SELECT id, title, description, status FROM tasks WHERE id = ?", id).
		Scan(&task.ID, &task.Title, &task.Description, &task.Status)
	if err != nil {
		return models.Task{}, err
	}

	return task, nil
}

func CreateTask(task models.Task) (int, error) {
	result, err := DB.Exec("INSERT INTO tasks (title, description, status) VALUES (?, ?, ?)",
		task.Title, task.Description, task.Status)
	if err != nil {
		return 0, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	return int(id), nil
}

func UpdateTask(task models.Task) error {
	_, err := DB.Exec("UPDATE tasks SET title = ?, description = ?, status = ? WHERE id = ?",
		task.Title, task.Description, task.Status, task.ID)
	return err
}

func DeleteTask(id int) error {
	_, err := DB.Exec("DELETE FROM tasks WHERE id = ?", id)
	return err
}
