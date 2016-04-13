package core

import (
	"html/template"
	"net/http"
	"strconv"

	"github.com/gorilla/context"
	"github.com/gorilla/mux"
	"github.com/gorilla/sessions"
	"gopkg.in/mgo.v2"

	"github.com/xjdrew/goweb/app/utils"
)

type Controller struct {
}

func (controller *Controller) GetSession(r *http.Request) *sessions.Session {
	return context.Get(r, sessionKey).(*sessions.Session)
}

func (controller *Controller) GetTemplate(r *http.Request) *template.Template {
	return context.Get(r, templateKey).(*template.Template)
}

func (controller *Controller) GetDatabase(r *http.Request) *mgo.Database {
	return context.Get(r, databaseKey).(*mgo.Database)
}

func (controller *Controller) IsPost(r *http.Request) bool {
	return r.Method == "POST"
}

func (controller *Controller) GetVar(r *http.Request, name string) (string, bool) {
	vars := mux.Vars(r)
	v, ok := vars[name]
	return v, ok
}

func (controller *Controller) GetVarInt(r *http.Request, name string) int {
	v, ok := controller.GetVar(r, name)
	if !ok {
		return 0
	}

	i, _ := strconv.Atoi(v)
	return i
}

func (controller *Controller) ReturnTemplate(r *http.Request, template string, c *C) (string, int) {
	t := controller.GetTemplate(r)
	content, err := utils.Parse(t, template, c.Env)
	if err != nil {
		return err.Error(), http.StatusInternalServerError
	}
	return content, http.StatusOK
}
