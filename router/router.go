package router

import "fmt"

func Test() {
	fmt.Println("Router package is linked")
}

/* type person struct {
    name string
    age  int
} */

type packet struct {
	Goal string
	Info string
}

func HandlePacket(packet packet) {
	// pull the goal

	// pull the info
	goal := packet.Goal

	switch goal {
	case "INIT":
		InitialConnection()
	case "REQLIST":
		RequestList()
	case "REQUSER":
		RequestUser()

	}

}

// Scan the additional information for username
func InitialConnection(packet packet) {
	return packet.Info
}
