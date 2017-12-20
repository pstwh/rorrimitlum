package main

import (
	"fmt"
	"encoding/gob"
	"net"
	"os"

	"github.com/go-vgo/robotgo"
)

type (
	Master struct {
		MouseX int
		MouseY int
	}

	Slave struct {
		Width  int
		Height int
	}
)

func main() {
	m := os.Args[1]

	//width, height := robotgo.GetScreenSize()

	conn, _ := net.Dial("tcp", m)
	defer conn.Close()

	decoder := gob.NewDecoder(conn)

	master := &Master{}

	for {
		decoder.Decode(master)
		fmt.Println(master)
		robotgo.Move(master.MouseX, master.MouseY)
	}

}
