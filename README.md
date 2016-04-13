Goweb
=====
goweb is a simple website based on [gorila toolkit](http://www.gorillatoolkit.org/) aiming to explain how to use gorilla toolkit.

it's similiar to [elcct's defaultproject](https://github.com/elcct/defaultproject), but don't depend on [goji](http://goji.io/). 

goweb show a simple way to use middleware and template with gorilla.

build
-----
```shell
# get code 
go get github.com/xjdrew/goweb

# get dependencies
go get github.com/gorilla/context
go get github.com/gorilla/mux
go get github.com/gorilla/schema
go get github.com/gorilla/sessions
go get gopkg.in/mgo.v2

# build
go install github.com/xjdrew/goweb/app
```

