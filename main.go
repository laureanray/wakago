package main

import (
	"os"
	"wakago/api"
	//	"wakago/cmd"
)

func main() {
	//	cmd.Execute()
	wt := api.GetInstance()

	wt.Init(os.Getenv("CLIENT_ID"), os.Getenv("CLIENT_SECRET"))
	wt.Login()
	api.InitServer()
}
