package main

import (
	"discobot/Handler"
	"discobot/Connection"
)

func main() {
	dg := Connection.Start()
	Handler.Register(dg)
	Connection.Wait(dg)
}
