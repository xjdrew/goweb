package models

import (
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type Server struct {
	ID           bson.ObjectId `bson:"_id,omitempty"`
	Serverid     int           `bson:"serverid"`
	Name         string        `bson:"name"`
	Status       int           `bson:"status"`
	Hidden       bool          `bson:"hidden"`
	Forbidden    bool          `bson:"forbidden"`
	MaxOnline    int           `bson:"max_online"`
	Default      int           `bson:"default"`
	RealServerid int           `bson:"real_serverid"`
	Device       int           `bson:"device"`
	Platform     string        `bson:"platform"`
}

func LoadServers(db *mgo.Database) (servers []Server, err error) {
	err = db.C("server").Find(nil).All(&servers)
	return
}

func InsertServer(db *mgo.Database, server *Server) error {
	server.ID = bson.NewObjectId()
	return db.C("server").Insert(server)
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
			"Platform":      server.Platform,
		},
	}
	return db.C("server").UpdateId(server.ID, change)
}
