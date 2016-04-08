package controllers

import (
	"net/http"

	"tea.ejoy.com/LR/smg/app/core"
	"tea.ejoy.com/LR/smg/app/utils"
)

type HomeController struct {
	core.Controller
}

func (controller *HomeController) Index(c *core.C, r *http.Request) (string, int) {
	t := controller.GetTemplate(r)

	c.Env["Title"] = "Home"
	content, err := utils.Parse(t, "home", c.Env)
	if err != nil {
		return err.Error(), http.StatusInternalServerError
	}
	return content, http.StatusOK
}
