package main

import (
	"encoding/gob"
	"fmt"
	"net"
	"os"
	"time"

	"./xrandr"
	"github.com/go-vgo/robotgo"
)

type (
	Master struct {
		MouseX int
		MouseY int
	}

	Slave struct {
		X int
		Y int
	}
)

var slaves = 0

func main() {
	xrandr.DisconnectAll()

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

		fmt.Println(*master)
		time.Sleep(10 * time.Microsecond)
	}
}

func slave(master *Master, conn net.Conn) {
	encoder := gob.NewEncoder(conn)

	var width, _ = robotgo.GetScreenSize()

	//fmt.Println(monitor)
	slaves++
	slaveX := width * slaves

	monitor := xrandr.Connect()

	for {
		if master.MouseX > slaveX+1366 {
			aux := Master{MouseX: 1366, MouseY: master.MouseY}
			encoder.Encode(aux)
		} else if master.MouseX < slaveX {
			aux := Master{MouseX: 0, MouseY: master.MouseY}
			encoder.Encode(aux)
		} else {
			aux := Master{MouseX: master.MouseX - slaveX,
				MouseY: master.MouseY}
			encoder.Encode(aux)
		}
		time.Sleep(1 * time.Millisecond)
	}

	xrandr.Disconnect(monitor)
	fmt.Println("Deleting Slave")
}
