package main

import (
	"encoding/json"
	"log"
	"net"
)

type Packet struct {
	Username   string
	Goal       string
	Info       string
	TargetUser string
	Address    net.Addr
}

func serializePacket(packet Packet) []byte {
	data, err := json.Marshal(packet)
	if err != nil {
		log.Fatal(err)
	}
	return data
}

func deserializePacket(data []byte) Packet {
	var packet Packet
	err := json.Unmarshal(data, &packet)
	if err != nil {
		log.Fatal(err)
	}

	return packet
}
