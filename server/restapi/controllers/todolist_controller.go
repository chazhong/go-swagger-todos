package controllers

import (
	"sync"
	"sync/atomic"

	errors "github.com/go-openapi/errors"

	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/swag"
	"github.com/chazhong/go-swagger-todos/server/models"
	"github.com/chazhong/go-swagger-todos/server/restapi/operations/todos"
)

var todoList = make(map[int64]*models.Item)
var lock = &sync.Mutex{}
var lastID int64

func addItem(item *models.Item) error {
	if item == nil {
		return errors.New(500, "item must be present")
	}
	lock.Lock()
	defer lock.Unlock()
	newID := atomic.AddInt64(&lastID, 1)
	item.ID = newID
	todoList[newID] = item
	return nil
}

type TodoListController struct {
}

func (c *TodoListController) GetTodoList(params todos.GetTodosParams) middleware.Responder {
	results := []*models.Item{}
	for _, v := range todoList {
		results = append(results, v)
	}
	return todos.NewGetTodosOK().WithPayload(results)
}

func (c *TodoListController) AddTodoItem(params todos.AddOneParams) middleware.Responder {
	if err := addItem(params.Body); err != nil {
		return todos.NewAddOneDefault(500).WithPayload(&models.Error{Code: 500, Message: swag.String(err.Error())})
	}
	return todos.NewAddOneCreated().WithPayload(params.Body)
}
