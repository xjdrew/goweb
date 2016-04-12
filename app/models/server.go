package models

import (
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

const (
	ServerStatusNew = 1
	ServerStatusHot = 2
)

type Server struct {
	ID           bson.ObjectId `bson:"_id,omitempty" schema:"-"`
	Serverid     int           `bson:"serverid"`
	Name         string        `bson:"name"`
	Status       int           `bson:"status"`
	Hidden       bool          `bson:"hidden"`
	Forbidden    bool          `bson:"forbidden"`
	MaxOnline    int           `bson:"max_online"`
	Default      int           `bson:"default"`
	RealServerid int           `bson:"real_serverid"`
	Device       int           `bson:"device"`
	Platform     int           `bson:"platform"`
}

func (server *Server) IsNew() bool {
	return server.Status&ServerStatusNew != 0
}

func (server *Server) IsHot() bool {
	return server.Status&ServerStatusHot != 0
}

func LoadServers(db *mgo.Database) (servers []Server, err error) {
	err = db.C("servers").Find(nil).Sort("serverid").All(&servers)
	return
}

func GetServer(db *mgo.Database, serverid int) (server *Server, err error) {
	server = new(Server)
	err = db.C("servers").Find(bson.M{"serverid": serverid}).One(server)
	if err != nil {
		return
	}
	return
}

func AllocServerId(db *mgo.Database) (err error, serverid int) {
	server := new(Server)
	err = db.C("servers").Find(nil).Sort("-serverid").One(server)
	if err != nil && err != mgo.ErrNotFound {
		return
	}

	if err == mgo.ErrNotFound {
		err = nil
		serverid = 1
	} else {
		serverid = server.Serverid + 1
	}
	return
}

func InsertServer(db *mgo.Database, server *Server) error {
	server.ID = bson.NewObjectId()
	return db.C("servers").Insert(server)
}

func UpdateServer(db *mgo.Database, server *Server) error {
	change := bson.M{
		"$set": bson.M{
			"name":          server.Name,
			"status":        server.Status,
			"hidden":        server.Hidden,
			"forbidden":     server.Forbidden,
			"max_online":    server.MaxOnline,
			"default":       server.Default,
			"real_serverid": server.RealServerid,
			"device":        server.Device,
			"platform":      server.Platform,
		},
	}
	return db.C("servers").Update(bson.M{"serverid": server.Serverid}, change)
}

func NewServer() *Server {
	server := new(Server)
	server.Status = 1
	server.MaxOnline = 10000
	return server
}
