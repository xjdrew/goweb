package main

import (
	"flag"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"

	"github.com/xjdrew/goweb/app/controllers"
	"github.com/xjdrew/goweb/app/core"
	"github.com/xjdrew/goweb/app/utils"
)

func usage() {
	log.Printf("Usage: %s [options] config\n", os.Args[0])
	flag.PrintDefaults()
	os.Exit(1)
}

func main() {
	listen := flag.String("listen", ":8080", "http server listen address")
	flag.Usage = usage
	flag.Parse()
	args := flag.Args()
	if len(args) == 0 {
		usage()
		return
	}

	app, err := core.NewApplication(args[0])
	if err != nil {
		log.Fatalf("new app failed:%v", err)
	}
	app.Init()
	defer app.Fini()

	r := mux.NewRouter()
	r.StrictSlash(true)
	hc := &controllers.HomeController{}
	r.HandleFunc("/", app.WrapRoute(hc.Index))

	s := r.PathPrefix("/server").Subrouter()
	sc := &controllers.ServerController{}
	s.HandleFunc("/", app.WrapRoute(sc.List))
	s.HandleFunc("/add", app.WrapRoute(sc.Add))
	s.HandleFunc("/update/{id:[0-9]+}", app.WrapRoute(sc.Update))

	r.PathPrefix("/assets/").Handler(http.StripPrefix("/assets", http.FileServer(http.Dir(app.Settings.PublicPath))))

	h := utils.UseMiddleware(r, app.ApplySession, app.ApplyTemplate, app.ApplyDatabase)
	log.Println(http.ListenAndServe(*listen, h))
}
