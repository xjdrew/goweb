package core

import (
	"net/http"

	"github.com/gorilla/context"
)

func (app *Application) ApplySession(h http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		session, _ := app.Store.Get(r, "session")
		context.Set(r, sessionKey, session)
		h.ServeHTTP(w, r)
	}
	return http.HandlerFunc(fn)
}

func (app *Application) ApplyTemplate(h http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		context.Set(r, templateKey, app.Template)
		h.ServeHTTP(w, r)
	}
	return http.HandlerFunc(fn)
}

func (app *Application) ApplyDatabase(h http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		context.Set(r, databaseKey, app.DBSession.DB(app.Settings.Database.Database))
		h.ServeHTTP(w, r)
	}
	return http.HandlerFunc(fn)
}
