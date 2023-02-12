package tasks

import (
	"fmt"
	"log"

	"awesomeProject/cmd/main/config"

	"github.com/jmoiron/sqlx"
)

func AddTaskDatabase(tasks []Task, db *sqlx.DB) error {
	for _, task := range tasks {
		sql := "INSERT INTO tasks (title, description, author) VALUES (:title, :description, :author)"
		result, err := db.NamedExec(sql, task)
		if err != nil {
			return err
		}
		count, _ := result.RowsAffected()
		log.Println(config.Yellow, fmt.Sprintf("Добавлено в базу данных строк: %v", count), config.Reset)
	}
	return nil
}

func GetTask(uuid int, db *sqlx.DB) (Task, error) {
	var task Task
	err := db.Get(&task, "SELECT * FROM tasks WHERE id= $1", uuid)

	return task, err
}

func GetTasks(db *sqlx.DB) ([]Task, error) {
	var tasks []Task
	err := db.Select(&tasks, "SELECT * FROM tasks")

	return tasks, err
}
