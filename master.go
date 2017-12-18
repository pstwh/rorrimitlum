package main

import (
	"encoding/gob"
	"fmt"
	"net"
	"os"
	"time"

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
)

func main() {
	var master = Master{}
	m := os.Args[1]

	fmt.Println("Master on:", m)
	ln, _ := net.Listen("tcp", m)

	go mouse(&master)

	for {
		conn, _ := ln.Accept()
		defer conn.Close()

		go slave(&master, conn)
	}

}

func mouse(master *Master) {
	for {
		master.MouseX, master.MouseY = robotgo.GetMousePos()
		time.Sleep(1 * time.Microsecond)
	}
}

func slave(master *Master, conn net.Conn) {
	fmt.Println("Adding Slave")
	encoder := gob.NewEncoder(conn)

	for {
		encoder.Encode(master)
		time.Sleep(10 * time.Millisecond)
	}
	fmt.Println("Deleting Slave")
}
