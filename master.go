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

var master = Master{}

func main() {
	fmt.Println("Master speaking")

	conn, _ := net.Dial("tcp", "127.0.0.1:8081")
	encoder := gob.NewEncoder(conn)

	for {
		mouseX, mouseY := robotgo.GetMousePos()

		master.MouseX = mouseX
		master.MouseY = mouseY

		encoder.Encode(master)
	}

	conn.Close()
}
