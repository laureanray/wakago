package api

import (
	"log"
	"net/http"
	"text/template"

	"github.com/julienschmidt/httprouter"
)

func Callback(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	w.Header().Set("Content-Type", "text/html; charset=utf8")
	t, err := template.ParseFiles("static/callback.html")
	if err != nil {
		log.Println(err)
	}
	code := r.URL.Query().Get("code")
	wt := GetInstance()
	log.Println(wt.oauthToken, code)

	wt.Exchange(code)

	if err != nil {
		log.Fatal(err)
	}

	t.Execute(w, nil)
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

func Kill() {

}
