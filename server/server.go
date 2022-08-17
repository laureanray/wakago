package server

import (
	"fmt"
	"log"
	"net/http"
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
	code := r.URL.Query().Get("code")
	wt := api.GetInstance()
	wt.Exchange(code)
	fmt.Fprint(w, "Welcome!\n")

}

func (s *Server) Init() {
	log.Println("Waiting for redirect URI")
	(*s).router = httprouter.New()
	(*s).server = &http.Server{Addr: ":8090", Handler: (*s).router}
	(*s).router.GET("/wakago/callback", loginCallback)
	log.Println("Server started")
	log.Fatal((*s).server.ListenAndServe())
}
