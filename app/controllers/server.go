package controllers

import (
	"log"
	"net/http"

	"github.com/gorilla/schema"

	"github.com/xjdrew/goweb/app/core"
	"github.com/xjdrew/goweb/app/models"
)

type serverKey int

const (
	addServerKey = iota
	updateServerKey
)

const (
	addServerFlashKey    = "AddServer"
	updateServerFlashKey = "UpdateServer"
)

type ServerController struct {
	core.Controller
}

func (controller *ServerController) List(c *core.C, r *http.Request) (string, int) {
	db := controller.GetDatabase(r)
	servers, err := models.LoadServers(db)
	if err != nil {
		return err.Error(), http.StatusInternalServerError
	}

	c.Env["Title"] = "列表"
	c.Env["Servers"] = servers
	return controller.ReturnTemplate(r, "server.list", c)
}

func (controller *ServerController) Add(c *core.C, r *http.Request) (string, int) {
	c.Env["Title"] = "添加"

	session := controller.GetSession(r)
	if controller.IsPost(r) {
		if err := r.ParseForm(); err != nil {
			session.AddFlash(err.Error(), addServerFlashKey)
			return r.URL.String(), http.StatusSeeOther
		}

		server := models.NewServer()
		decoder := schema.NewDecoder()
		if err := decoder.Decode(server, r.PostForm); err != nil {
			session.Values[addServerKey] = server
			session.AddFlash(err.Error(), addServerFlashKey)
			return r.URL.String(), http.StatusSeeOther
		}

		log.Printf("add: %+v", server)

		db := controller.GetDatabase(r)
		err, serverid := models.AllocServerId(db)
		if err != nil {
			session.Values[addServerKey] = server
			session.AddFlash(err.Error(), addServerFlashKey)
			return r.URL.String(), http.StatusSeeOther
		}
		server.Serverid = serverid

		err = models.InsertServer(db, server)
		if err != nil {
			session.Values[addServerKey] = server
			session.AddFlash(err.Error(), addServerFlashKey)
			return r.URL.String(), http.StatusSeeOther
		}
		delete(session.Values, addServerKey)

		c.Env["Serverid"] = serverid
		return controller.ReturnTemplate(r, "server.redirect", c)
	} else {
		c.Env["Flash"] = session.Flashes(addServerFlashKey)

		if server, ok := session.Values[addServerKey]; ok {
			c.Env["Server"] = server
		} else {
			c.Env["Server"] = models.NewServer()
		}

		return controller.ReturnTemplate(r, "server.add", c)
	}
}

func (controller *ServerController) Update(c *core.C, r *http.Request) (string, int) {
	c.Env["Title"] = "更新"

	session := controller.GetSession(r)
	db := controller.GetDatabase(r)

	if controller.IsPost(r) {
		if err := r.ParseForm(); err != nil {
			session.AddFlash(err.Error(), updateServerFlashKey)
			return r.URL.String(), http.StatusSeeOther
		}

		server := models.NewServer()
		decoder := schema.NewDecoder()
		if err := decoder.Decode(server, r.PostForm); err != nil {
			session.Values[updateServerKey] = server
			session.AddFlash(err.Error(), updateServerFlashKey)
			return r.URL.String(), http.StatusSeeOther
		}

		log.Printf("update: %+v", server)

		if err := models.UpdateServer(db, server); err != nil {
			session.Values[updateServerKey] = server
			session.AddFlash(err.Error(), updateServerFlashKey)
			return r.URL.String(), http.StatusSeeOther
		}
		delete(session.Values, updateServerKey)

		c.Env["Serverid"] = server.Serverid
		return controller.ReturnTemplate(r, "server.redirect", c)
	} else {
		serverid := controller.GetVarInt(r, "id")
		if old, ok := session.Values[updateServerKey].(*models.Server); ok && old.Serverid == serverid {
			c.Env["Server"] = old
		} else {
			server, err := models.GetServer(db, serverid)
			if err != nil {
				session.AddFlash(err.Error(), updateServerFlashKey)
			} else {
				c.Env["Server"] = server
			}
		}
		c.Env["Flash"] = session.Flashes(updateServerFlashKey)

		return controller.ReturnTemplate(r, "server.update", c)
	}
}
