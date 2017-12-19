package main

import (
	"encoding/gob"
	"fmt"
	"net"
	"os"
	"time"

	"github.com/go-vgo/robotgo"
	"rorrimitlum/xrandr"
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

var slaves = 1
var width, height = robotgo.GetScreenSize()

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

		fmt.Println(*master)
		time.Sleep(10 * time.Microsecond)
	}
}

func slave(master *Master, conn net.Conn) {
	encoder := gob.NewEncoder(conn)

	monitor := xrandr.Connect()
	//fmt.Println(monitor)
	slaves++
	slaveX := width * slaves

	for {
		if master.MouseX > slaveX+1366 {
			aux := Master{MouseX: slaveX + 1366, MouseY: master.MouseY}
			encoder.Encode(aux)
		} else if master.MouseX < slaveX {
			aux := Master{MouseX: slaveX, MouseY: master.MouseY}
			encoder.Encode(aux)
		} else {
			encoder.Encode(master)
		}
		time.Sleep(1 * time.Millisecond)
	}

	xrandr.Disconnect(monitor)
	fmt.Println("Deleting Slave")
}
