package server

import (
	"context"
	"log"
	"net/http"
	"text/template"
	"wakago/api"

	"github.com/julienschmidt/httprouter"
)

type Server struct {
	router *httprouter.Router
	server *http.Server
}

var serverInstance *Server

func GetInstance() *Server {
	if serverInstance == nil {
		serverInstance = new(Server)
	}

	return serverInstance
}

func loginCallback(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	w.Header().Set("Content-Type", "text/html; charset=utf8")
	t, err := template.ParseFiles("static/callback.html")
	if err != nil {
		log.Println(err)
	}
	code := r.URL.Query().Get("code")
	wt := api.GetInstance()
	wt.Exchange(code)

	if err != nil {
		log.Fatal(err)
	}

	t.Execute(w, nil)
	if err != nil {
		log.Println("e", err)
	}

	serverInstance.server.Shutdown(context.TODO())
}

func (s *Server) Init() {
	log.Println("Waiting for redirect URI")
	(*s).router = httprouter.New()
	(*s).server = &http.Server{Addr: ":8090", Handler: (*s).router}
	(*s).router.GET("/wakago/callback", loginCallback)
	log.Println("Server started")
	log.Fatal((*s).server.ListenAndServe())
}
