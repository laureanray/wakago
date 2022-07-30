package main

import (
	"wakago/api"
	"wakago/cmd"
)

func main() {
	cmd.Execute()
	api.GetInstance()
}
