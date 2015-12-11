package main

import "system"

func main() {
	server := system.NewServer("tcp4", ":9999", nil, nil)
	server.Run()
}
