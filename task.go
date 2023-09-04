package main

type Task struct {
	Id          int    `json:"id"`
	Discription string `json:"discription"`
	IsDone      bool   `json:"isDone"`
}
