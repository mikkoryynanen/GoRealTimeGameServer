package main

import (
	network "github.com/mikkoryynanen/real-time/internal/network"
	arguments "github.com/mikkoryynanen/real-time/utils"
)

func main() {
	args := arguments.GetArguments()
	network.Start(args)
}