package models

import "go-crud/db"

func GetAll() (todos []Todo, err error) {
	conn, err := db.OpenConnection()
	if err != nil {
		return
	}
	defer conn.Close()

	sql := "SELECT id, title, description, done FROM todos"
	rows, err := conn.Query(sql)
	if err != nil {
		return
	}

	for rows.Next() {
		var todo Todo

		err = conn.QueryRow(sql).Scan(&todo.Id, &todo.Title, &todo.Description, &todo.Done)
		if err != nil {
			continue
		}

		todos = append(todos, todo)
	}

	return
}
