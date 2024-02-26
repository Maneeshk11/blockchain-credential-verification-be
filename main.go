package main

import "blockchain-verification-be/server"

func main() {
	serverMain := server.NewServer()
	serverMain.Run()
}
