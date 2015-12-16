//Copyright 2015 Libing. All right reserved.
package main

import "system/communicate"

func main() {
	c := communicate.NewClient("tcp", "localhost:9999")
	c.SendStrings(1, "Hello world, hi li", "Hi ok", "{\"name\":\"hello\"}")
	c.SendStrings(2, "this is example1, hi bli", "Hi ok1", "{\"name\":\"world1\"}")
	c.SendStrings(3, "this is example2, hi bli", "Hi ok2", "{\"name\":\"world2\"}")
	c.SendStrings(4, "this is example3, hi bli", "Hi ok3", "{\"name\":\"world3\"}")
	c.SendStrings(5, "this is example4, hi bli", "Hi ok4", "{\"name\":\"world4\"}")
	c.Close()
	c.ReConnect()
	c.SendStrings(6, "this is example6, hi bli", "Hi ok4", "{\"name\":\"world6\"}")

}
