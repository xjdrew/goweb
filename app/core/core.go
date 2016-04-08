package core

import (
	"html/template"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"github.com/gorilla/context"
	"github.com/gorilla/mux"
	"github.com/gorilla/sessions"

	"gopkg.in/mgo.v2"
	//"gopkg.in/mgo.v2/bson"
)

type contextKey int

const (
	appKey contextKey = iota
	sessionKey
	templateKey
)

type Application struct {
	Settings  *Settings
	Template  *template.Template
	Router    *mux.Router
	Store     *sessions.CookieStore
	DBSession *mgo.Session
}

func (app *Application) loadTemplates() {
	var templates []string
	fn := func(path string, f os.FileInfo, err error) error {
		if err != nil {
			log.Printf("failed load templates:%v", err)
			return nil
		}

		if !f.IsDir() && strings.HasSuffix(f.Name(), ".html") {
			templates = append(templates, path)
		}
		return nil
	}

	filepath.Walk(app.Settings.TemplatePath, fn)
	if len(templates) > 0 {
		app.Template = template.Must(template.ParseFiles(templates...))
	} else {
		app.Template = template.New("")
	}
}

func (app *Application) initDatabase() {
	session, err := mgo.Dial(app.Settings.Database.Hosts)
	if err != nil {
		log.Fatalf("can't connect to the database:%v", err)
	}
	app.DBSession = session
}

func (app *Application) Init() {
	app.loadTemplates()
	app.initDatabase()
	app.Store = sessions.NewCookieStore([]byte(app.Settings.Secret))
	app.Router = mux.NewRouter()
}

func (app *Application) Fini() {
	if app.DBSession != nil {
		app.DBSession.Close()
	}
}

func (app *Application) WrapRoute(f func(*C, *http.Request) (string, int)) http.HandlerFunc {
	fn := func(w http.ResponseWriter, r *http.Request) {
		context.Set(r, appKey, app)

		c := NewC()
		body, code := f(c, r)

		if session, ok := context.GetOk(r, sessionKey); ok {
			s := session.(*sessions.Session)
			s.Save(r, w)
		}

		switch code {
		case http.StatusOK:
			if ct, ok := c.Env["Content-Type"]; ok {
				w.Header().Set("Content-Type", ct.(string))
			} else {
				w.Header().Set("Content-Type", "text/html")
			}
			io.WriteString(w, body)
		case http.StatusSeeOther, http.StatusFound:
			http.Redirect(w, r, body, code)
		default:
			w.WriteHeader(code)
			io.WriteString(w, body)
		}
	}
	return fn
}

func NewApplication(file string) (*Application, error) {
	settings := &Settings{}
	if err := settings.Load(file); err != nil {
		return nil, err
	}
	app := &Application{
		Settings: settings,
	}
	return app, nil
}
