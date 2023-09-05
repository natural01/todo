package pkg

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type Handler struct {
	Db Database
}

func (h *Handler) NewHandler() *gin.Engine {
	h.Db.Run()
	router := gin.Default()
	router.GET("/tasks", h.getTasks)
	router.POST("/tasks", h.postTask)
	router.POST("/tasks/complete/:id", h.completeTask)
	router.GET("/tasks/completed", h.getCompletedTasks)
	router.GET("/tasks/uncompleted", h.getUncompletedTasks)
	router.GET("/tasks/:id", h.getTaskById)
	router.DELETE("/tasks/delete/:id", h.deleteTask)

	return router
}

func (h *Handler) getTasks(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, h.Db.GetTasks())
}

func (h *Handler) getCompletedTasks(c *gin.Context) {
	var completedTasks []Task
	for _, task := range h.Db.GetTasks() {
		if task.IsDone == true {
			completedTasks = append(completedTasks, task)
		}
	}
	c.IndentedJSON(http.StatusOK, completedTasks)
}

func (h *Handler) getUncompletedTasks(c *gin.Context) {
	var completedTasks []Task
	for _, task := range h.Db.GetTasks() {
		if task.IsDone == false {
			completedTasks = append(completedTasks, task)
		}
	}
	c.IndentedJSON(http.StatusOK, completedTasks)
}

func (h *Handler) postTask(c *gin.Context) {
	var newTask Task

	err := c.BindJSON(&newTask)
	if err != nil || !h.isValidParams(newTask) {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "task can not be added"})
		return
	}

	newId := h.Db.GetTasks()[h.getCountOfTasks()-1].Id
	newTask.Id = newId + 1
	h.Db.AddTask(newTask)
	c.IndentedJSON(http.StatusCreated, newTask)
}

func (h *Handler) completeTask(c *gin.Context) {
	id := c.Param("id")

	for _, taskId := range h.Db.GetTasks() {
		if strconv.Itoa(taskId.Id) == id {
			i, err := strconv.Atoi(id)
			if err != nil {
				panic(err)
			}
			h.Db.UpdateTaskState(i, true)
			c.IndentedJSON(http.StatusOK, taskId)
			return
		}
	}

	c.IndentedJSON(http.StatusNotAcceptable, gin.H{"message": "task can not be found"})
}

func (h *Handler) getTaskById(c *gin.Context) {
	id := c.Param("id")

	for _, taskId := range h.Db.GetTasks() {
		if strconv.Itoa(taskId.Id) == id {
			c.IndentedJSON(http.StatusOK, taskId)
			return
		}
	}

	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "task can not be found"})
}

func (h *Handler) deleteTask(c *gin.Context) {
	id := c.Param("id")

	for _, taskId := range h.Db.GetTasks() {
		if strconv.Itoa(taskId.Id) == id {
			i, err := strconv.Atoi(id)
			if err != nil {
				panic(err)
			}
			h.Db.DeleteTask(i)
			c.IndentedJSON(http.StatusOK, taskId)
			return
		}
	}

	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "task can not be found"})
}

func (h *Handler) getCountOfTasks() int {
	tasks := h.Db.GetTasks()
	var count = 0

	for _, _ = range tasks {
		count++
	}

	return count
}

func (h *Handler) isValidParams(task Task) bool {
	if task.Id != 0 || task.IsDone || len(task.Discription) > 1000 {
		return false
	}
	return true
}
