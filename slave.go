package main

import (
	"encoding/gob"
	"fmt"
	"net"

	"github.com/go-vgo/robotgo"
)

type Master struct {
	MouseX int
	MouseY int

	PressedLeftButton  bool
	PressedRightButton bool
	//PressedKeys []int
}

func main() {
	fmt.Println("Slave listening")

	ln, _ := net.Listen("tcp", "127.0.0.1:8081")

	conn, _ := ln.Accept()
	decoder := gob.NewDecoder(conn)

	master := &Master{}

	for {
		decoder.Decode(master)
		robotgo.MoveMouse(master.MouseX, master.MouseY)
	}
}
