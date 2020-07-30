package main

import (
	"net/http"
	"time"

	"github.com/casbin/casbin"
	mongodbadapter "github.com/casbin/mongodb-adapter"
	"github.com/go-chi/chi"
)

var enforcer *casbin.SyncedEnforcer

func main() {
	a := mongodbadapter.NewAdapter("127.0.0.1:27017")
	//e := casbin.NewEnforcer("path/to/model.conf", a)
	enforcer = casbin.NewSyncedEnforcer("./model.conf", a)
	enforcer.StartAutoLoadPolicy(time.Minute * 5)

	r := chi.NewRouter()

	r.Get("/api/rule/validate", func(w http.ResponseWriter, r *http.Request) {

	})
	r.Route("/api", func(r chi.Router) {

	})
	// r.Get("", func(w http.ResponseWriter, r *http.Request) {

	// })
	http.ListenAndServe(":3000", r)

}
