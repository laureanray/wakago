package main

import (
	"os"
	"wakago/api"
	"wakago/cmd"
)

func main() {
	wt := api.GetInstance()
	wt.Init(os.Getenv("CLIENT_ID"), os.Getenv("CLIENT_SECRET"))
	cmd.Execute()
}
