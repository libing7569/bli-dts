//Copyright 2015 Libing. All right reserved.

package main

import "system/communicate"

func main() {
	server := communicate.NewServer("tcp4", ":9999", nil, nil)
	server.Run()
}
