package core

import (
	"html/template"
	"net/http"

	"github.com/gorilla/context"
	"github.com/gorilla/sessions"
	"gopkg.in/mgo.v2"
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
	return nil
}
