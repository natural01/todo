package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

type Database struct {
	db *sql.DB
}

func (database *Database) Run() {
	db, err := sql.Open("mysql", "root:root@tcp(127.0.0.1:8889)/todo")
	if err != nil {
		panic(err)
	}
	database.db = db
}

func (database *Database) close(db *sql.DB) {
	err := db.Close()
	if err != nil {
		return
	}
}

func (database *Database) GetTasks() []Task {
	var tasks []Task

	res, err := database.db.Query("SELECT * FROM `tasks`")
	if err != nil {
		panic(err)
	}

	for res.Next() {
		var task Task
		err = res.Scan(&task.Id, &task.Discription, &task.IsDone)
		if err != nil {
			panic(err)
		}

		tasks = append(tasks, task)
	}

	return tasks
}

func (database *Database) AddTask(task Task) {
	insert, err := database.db.Query(fmt.Sprintf("INSERT INTO `tasks` (`id`, `discription`, `isDone`) VALUES('%d', '%s', '%d')", task.Id, task.Discription, database.boolToInt(task.IsDone)))
	if err != nil {
		panic(err)
	}
	err = insert.Close()
	if err != nil {
		return
	}
}

func (database *Database) UpdateTaskState(id int, state bool) {
	insert, err := database.db.Query(fmt.Sprintf("UPDATE `tasks` SET `isDone` = %d WHERE `id` = %d", database.boolToInt(state), id))
	if err != nil {
		panic(err)
	}
	err = insert.Close()
	if err != nil {
		return
	}
}

func (database *Database) DeleteTask(id int) {
	insert, err := database.db.Query(fmt.Sprintf("DELETE FROM `tasks` WHERE `id` = %d", id))
	if err != nil {
		panic(err)
	}
	err = insert.Close()
	if err != nil {
		return
	}
}

func (database *Database) boolToInt(state bool) int8 {
	switch state {
	case true:
		return 1
	case false:
		return 0
	}
	return 0
}
