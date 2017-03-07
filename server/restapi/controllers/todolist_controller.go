package controllers

import (
	"sync"
	"sync/atomic"

	errors "github.com/go-openapi/errors"

	"github.com/chazhong/go-swagger-todos/server/models"
	"github.com/chazhong/go-swagger-todos/server/restapi/operations/todos"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/swag"
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

func deleteItem(id int64) error {
	lock.Lock()
	defer lock.Unlock()

	_, exists := todoList[id]
	if !exists {
		return errors.NotFound("not found: item %d", id)
	}

	delete(todoList, id)
	return nil
}

func updateItem(id int64, item *models.Item) error {
	if item == nil {
		return errors.New(500, "item must be present")
	}

	lock.Lock()
	defer lock.Unlock()

	_, exists := todoList[id]
	if !exists {
		return errors.NotFound("not found: item %d", id)
	}

	item.ID = id
	todoList[id] = item
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

func (c *TodoListController) DeleteTodoItem(params todos.DestroyOneParams) middleware.Responder {
	if err := deleteItem(params.ID); err != nil {
		return todos.NewDestroyOneDefault(500).WithPayload(&models.Error{Code: 500, Message: swag.String(err.Error())})
	}
	return todos.NewDestroyOneNoContent()
}

func (c *TodoListController) UpdateTodoItem(params todos.UpdateOneParams) middleware.Responder {
	if err := updateItem(params.ID, params.Body); err != nil {
		return todos.NewUpdateOneDefault(500).WithPayload(&models.Error{Code: 500, Message: swag.String(err.Error())})
	}
	return todos.NewUpdateOneOK().WithPayload(params.Body)
}
