package main

import (
	"ToDo/pkg"
	"encoding/json"
	. "github.com/smartystreets/goconvey/convey"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strconv"
	"strings"
	"testing"
)

func TestGetTasks(t *testing.T) {
	h := new(pkg.Handler).NewHandler()
	var id int

	Convey("Tasks should be found", t, func() {
		req, _ := http.NewRequest(http.MethodGet, "http://localhost:8080/tasks", nil)
		testRes := httptest.NewRecorder()
		h.ServeHTTP(testRes, req)

		So(testRes.Code, ShouldEqual, http.StatusOK)
	})

	Convey("Completed tasks should be found", t, func() {
		req, _ := http.NewRequest(http.MethodGet, "http://localhost:8080/tasks/completed", nil)
		testRes := httptest.NewRecorder()
		h.ServeHTTP(testRes, req)

		So(testRes.Code, ShouldEqual, http.StatusOK)
	})

	Convey("Uncompleted tasks should be found", t, func() {
		req, _ := http.NewRequest(http.MethodGet, "http://localhost:8080/tasks/uncompleted", nil)
		testRes := httptest.NewRecorder()
		h.ServeHTTP(testRes, req)

		So(testRes.Code, ShouldEqual, http.StatusOK)
	})

	Convey("Task should be added", t, func() {
		bodyReader := strings.NewReader(`{"discription": "12124"}`)
		req, _ := http.NewRequest(http.MethodPost, "http://localhost:8080/tasks", bodyReader)
		req.Header.Set("content-type", "application/json; charset=utf-8")
		testRes := httptest.NewRecorder()
		h.ServeHTTP(testRes, req)

		body, _ := ioutil.ReadAll(testRes.Body)
		var task pkg.Task
		err := json.Unmarshal(body, &task)
		if err != nil {
			panic(err)
		}
		id = task.Id

		So(testRes.Code, ShouldEqual, http.StatusCreated)
	})

	Convey("Task should be found", t, func() {
		req, _ := http.NewRequest(http.MethodGet, "http://localhost:8080/tasks/"+strconv.Itoa(id), nil)
		testRes := httptest.NewRecorder()
		h.ServeHTTP(testRes, req)

		So(testRes.Code, ShouldEqual, http.StatusOK)
	})

	Convey("Task should be updated", t, func() {
		req, _ := http.NewRequest(http.MethodPost, "http://localhost:8080/tasks/complete/"+strconv.Itoa(id), nil)
		req.Header.Set("content-type", "application/json; charset=utf-8")
		testRes := httptest.NewRecorder()
		h.ServeHTTP(testRes, req)

		So(testRes.Code, ShouldEqual, http.StatusOK)
	})

	Convey("Task should be deleted", t, func() {
		req, _ := http.NewRequest(http.MethodDelete, "http://localhost:8080/tasks/delete/"+strconv.Itoa(id), nil)
		req.Header.Set("content-type", "application/json; charset=utf-8")
		testRes := httptest.NewRecorder()
		h.ServeHTTP(testRes, req)

		So(testRes.Code, ShouldEqual, http.StatusOK)
	})

	Convey("Task shouldn't be found", t, func() {
		req, _ := http.NewRequest(http.MethodGet, "http://localhost:8080/tasks/"+strconv.Itoa(id), nil)
		testRes := httptest.NewRecorder()
		h.ServeHTTP(testRes, req)

		So(testRes.Code, ShouldEqual, http.StatusNotFound)
	})

	Convey("Task shouldn't be updated", t, func() {
		req, _ := http.NewRequest(http.MethodPost, "http://localhost:8080/tasks/complete/"+strconv.Itoa(id), nil)
		req.Header.Set("content-type", "application/json; charset=utf-8")
		testRes := httptest.NewRecorder()
		h.ServeHTTP(testRes, req)

		So(testRes.Code, ShouldEqual, http.StatusNotAcceptable)
	})

	Convey("Task shouldn't be deleted", t, func() {
		req, _ := http.NewRequest(http.MethodDelete, "http://localhost:8080/tasks/delete/"+strconv.Itoa(id), nil)
		req.Header.Set("content-type", "application/json; charset=utf-8")
		testRes := httptest.NewRecorder()
		h.ServeHTTP(testRes, req)

		So(testRes.Code, ShouldEqual, http.StatusNotFound)
	})

	Convey("Task shouldn't be added", t, func() {
		bodyReader := strings.NewReader(`{"id": 1, "discription": "12124", "isDone": true}`)
		req, _ := http.NewRequest(http.MethodPost, "http://localhost:8080/tasks", bodyReader)
		req.Header.Set("content-type", "application/json; charset=utf-8")
		testRes := httptest.NewRecorder()
		h.ServeHTTP(testRes, req)

		body, _ := ioutil.ReadAll(testRes.Body)
		var task pkg.Task
		err := json.Unmarshal(body, &task)
		if err != nil {
			panic(err)
		}
		id = task.Id

		So(testRes.Code, ShouldEqual, http.StatusNotFound)
	})
}
