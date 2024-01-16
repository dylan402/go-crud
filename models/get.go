package models

import "go-crud/db"

func Get(id int64) (todo Todo, err error) {
	conn, err := db.OpenConnection()
	if err != nil {
		return
	}
	defer conn.Close()

	sql := "SELECT id, title, description, done FROM todos WHERE id = $1"

	err = conn.QueryRow(sql, id).Scan(&todo.Id, &todo.Title, &todo.Description, &todo.Done)

	return
}
