package main

import "net"

func routePacket(packet Packet, conn net.Conn) {

	switch packet.Goal {
	case "initial_connection":
		initialConnection(packet)
	case "request_list":
		//sendUserList(conn)
	case "request_user":
		//sendUserRequest(packet)
	case "request_user_response":
		//processResponse(packet)
	default:
		//handleBadRequest()
	}

}
