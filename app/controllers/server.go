package controllers

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/schema"

	"tea.ejoy.com/LR/smg/app/core"
	"tea.ejoy.com/LR/smg/app/models"
	"tea.ejoy.com/LR/smg/app/utils"
)

type ServerController struct {
	core.Controller
}

func (controller *ServerController) List(c *core.C, r *http.Request) (string, int) {
	t := controller.GetTemplate(r)

	c.Env["Title"] = "列表"
	content, err := utils.Parse(t, "server_list", c.Env)
	if err != nil {
		return err.Error(), http.StatusInternalServerError
	}
	return content, http.StatusOK
}

func (controller *ServerController) Add(c *core.C, r *http.Request) (string, int) {
	session := controller.GetSession(r)

	if controller.IsPost(r) {
		if err := r.ParseForm(); err != nil {
			session.AddFlash(err.Error(), "AddServer")
			return r.URL.String(), http.StatusSeeOther
		}

		log.Printf("%+v", r.PostForm)

		server := models.NewServer()
		decoder := schema.NewDecoder()
		if err := decoder.Decode(server, r.PostForm); err != nil {
			session.Values["Server"] = server
			session.AddFlash(err.Error(), "AddServer")
			return r.URL.String(), http.StatusSeeOther
		}

		return fmt.Sprintf("%+v", server), http.StatusOK
	} else {
		t := controller.GetTemplate(r)

		c.Env["Title"] = "添加"
		c.Env["Flash"] = session.Flashes("AddServer")

		if server, ok := session.Values["Server"]; ok {
			c.Env["Server"] = server
		} else {
			c.Env["Server"] = models.NewServer()
		}

		content, err := utils.Parse(t, "server_add", c.Env)
		if err != nil {
			return err.Error(), http.StatusInternalServerError
		}
		return content, http.StatusOK
	}
}

func (controller *ServerController) Update(c *core.C, r *http.Request) (string, int) {
	t := controller.GetTemplate(r)
	c.Env["Title"] = "更新"
	content, err := utils.Parse(t, "server_update", c.Env)
	if err != nil {
		return err.Error(), http.StatusInternalServerError
	}
	return content, http.StatusOK
}
