package controllers

import (
	"net/http"

	"tea.ejoy.com/LR/smg/app/core"
)

type ServerController struct {
	core.Controller
}

func (controller *ServerController) List(r *http.Request) (string, int) {
	return "list", http.StatusOK
}

func (controller *ServerController) Add(r *http.Request) (string, int) {
	return "add", http.StatusOK
}

func (controller *ServerController) Update(r *http.Request) (string, int) {
	return "update", http.StatusOK
}

func (controller *ServerController) Delete(r *http.Request) (string, int) {
	return "delete", http.StatusOK
}
