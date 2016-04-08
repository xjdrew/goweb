package controllers

import (
	"net/http"

	"tea.ejoy.com/LR/smg/app/core"
	"tea.ejoy.com/LR/smg/app/utils"
)

type ServerController struct {
	core.Controller
}

func (controller *ServerController) List(c *core.C, r *http.Request) (string, int) {
	t := controller.GetTemplate(r)

	c.Env["Title"] = "server list"
	content, err := utils.Parse(t, "server_list", c.Env)
	if err != nil {
		return err.Error(), http.StatusInternalServerError
	}
	return content, http.StatusOK
}

func (controller *ServerController) Add(c *core.C, r *http.Request) (string, int) {
	return "add", http.StatusOK
}

func (controller *ServerController) Update(c *core.C, r *http.Request) (string, int) {
	return "update", http.StatusOK
}

func (controller *ServerController) Delete(c *core.C, r *http.Request) (string, int) {
	return "delete", http.StatusOK
}
