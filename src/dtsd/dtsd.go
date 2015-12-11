//Copyright 2015 Libing. All right reserved.

package main

import "system/server"

func main() {
	server := server.NewServer("tcp4", ":9999", nil, nil)
	server.Run()
}
