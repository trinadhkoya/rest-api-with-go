package controller

import (
	"fmt"
	"net/http"
)

type controller struct {

}

func (c controller) Welcome(w http.ResponseWriter, r *http.Request) {
	_, _ = fmt.Fprintf(w,"Hello! Welcome to go!")
}

func (c controller) GetHealth(w http.ResponseWriter, r *http.Request) {
	_, _ = fmt.Fprintf(w,"Health is OK")
}

type UtilsController interface {
	GetHealth(w http.ResponseWriter, r *http.Request)
	Welcome(w http.ResponseWriter, r *http.Request)
}

func NewUtilsController() UtilsController {
	return &controller{}
}