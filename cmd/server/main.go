package main

import (
	"log"

	network "github.com/mikkoryynanen/real-time/internal/network"
	arguments "github.com/mikkoryynanen/real-time/utils"
)

func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	
	args := arguments.GetArguments()
	log.Fatal(network.Start(args))
}