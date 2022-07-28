package api

import (
	"fmt"
	"log"
	"net/http"
	"text/template"

	"github.com/julienschmidt/httprouter"
)

func Callback(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	w.Header().Set("Content-Type", "text/html; charset=utf8")
	log.Println(r.URL.Query().Get("code"))
	t, err := template.ParseFiles("static/callback.html")

	if err != nil {
		log.Println(err)
	}

	fmt.Println(t)
	fmt.Println("Callback called!")
	//fmt.Fprint(w, t)

	err = t.Execute(w, nil)
	if err != nil {
		log.Println("e", err)
	}
}

func InitServer() {
	log.Println("Waiting for redirect URI")
	router := httprouter.New()

	router.GET("/wakago/callback", Callback)
	log.Println("Server started")

	log.Fatal(http.ListenAndServe(":8090", router))
}
