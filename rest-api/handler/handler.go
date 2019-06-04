package handler

import "github.com/lucfek/go-exercises/rest-api/model"

type modelInter interface {
	Set(todo model.Todo) (model.Todo, error)
	Get(id uint64) (model.Todo, error)
	GetAll() ([]model.Todo, error)
	Update(id uint64, todo model.Todo) (model.Todo, error)
	Delete(id uint64) (model.Todo, error)
}

type postData struct {
	Name string
	Desc string
}

// Handler is a struct responsible for handling requests
type Handler struct {
	m modelInter
}

// New is a constructor of "Handler", it gets "Model" type Model as an argument and returns "Handler" type Handler
func New(m modelInter) Handler {
	return Handler{m: m}
}
