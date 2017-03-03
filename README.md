# go-swagger-todos
Todolist server generated using go-swagger

## Generate Server
```
swagger generate server -A TodoList -t server -f ./swagger.yml
```

## Create Controllers
```
mkdir server/restapi/controllers
touch server/restapi/controllers/todolist_controller.go
```
```javascript
// todolist_controller.go
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

```

## Set API Handlers
```javascript
// configure_todo_list.go
func configureAPI(api *operations.TodoListAPI) http.Handler {
    ...
    // Controllers
    todoListController := controllers.TodoListController{}

    // API Handlers
    api.TodosGetTodosHandler = todos.GetTodosHandlerFunc(todoListController.GetTodoList)
    api.TodosAddOneHandler = todos.AddOneHandlerFunc(todoListController.AddTodoItem)
    ...
}
```