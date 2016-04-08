package main

import (
	"flag"
	"log"
	"net/http"
	"os"

	"tea.ejoy.com/LR/smg/app/controllers"
	"tea.ejoy.com/LR/smg/app/core"
	"tea.ejoy.com/LR/smg/app/utils"
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

	r := app.Router
	hc := &controllers.HomeController{}
	r.HandleFunc("/", app.WrapRoute(hc.Index))

	s := r.PathPrefix("/server").Subrouter()
	sc := &controllers.ServerController{}
	s.HandleFunc("/", app.WrapRoute(sc.List))
	s.HandleFunc("/add", app.WrapRoute(sc.Add))
	s.HandleFunc("/update", app.WrapRoute(sc.Update))
	s.HandleFunc("/delete", app.WrapRoute(sc.Delete))

	r.PathPrefix("/assets/").Handler(http.StripPrefix("/assets", http.FileServer(http.Dir(app.Settings.PublicPath))))

	h := utils.UseMiddleware(r, app.ApplySession, app.ApplyTemplate)
	log.Println(http.ListenAndServe(*listen, h))
}
