package main

import (
	"log"
	"os"
	"wakago/api"
	//	"wakago/cmd"
)

func main() {
	//	cmd.Execute()
	wt := api.GetInstance()

	log.Println("CLIENT_ID", os.Getenv("CLIENT_ID"))
	log.Println("CLIENT_SECRET", os.Getenv("CLIENT_SECRET"))

	wt.Init(os.Getenv("CLIENT_ID"), os.Getenv("CLIENT_SECRET"))
	wt.Login()
	api.InitServer()
}
