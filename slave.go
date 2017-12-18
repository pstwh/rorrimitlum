package main

import (
	"encoding/gob"
	"fmt"
	"net"

	"os"

	"github.com/go-vgo/robotgo"
)

type (
	Master struct {
		MouseX int
		MouseY int

		PressedLeftButton  bool
		PressedRightButton bool
		//PressedKeys []int
	}

	Slave struct {
		Width  int
		Height int
	}
)

func main() {
	m := os.Args[1]

	width, height := robotgo.GetScreenSize()
	fmt.Println("W:", width, "H:", height)

	conn, _ := net.Dial("tcp", m)
	defer conn.Close()

	decoder := gob.NewDecoder(conn)

	master := &Master{}

	for {
		decoder.Decode(master)
		robotgo.Move(master.MouseX, master.MouseY)
	}

}
