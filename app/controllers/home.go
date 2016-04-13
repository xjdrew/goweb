package controllers

import (
	"net/http"

	"github.com/xjdrew/goweb/app/core"
)

type HomeController struct {
	core.Controller
}

func (controller *HomeController) Index(c *core.C, r *http.Request) (string, int) {
	c.Env["Title"] = "Home"
	return controller.ReturnTemplate(r, "home", c)
}
