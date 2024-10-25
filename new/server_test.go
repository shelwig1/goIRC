package main

import (
	"net"
	"testing"
	"time"
)

type MockConn struct {
	net.Conn
	sendData []byte
}

func (m *MockConn) Write(b []byte) (int, error) {
	m.sendData = append(m.sendData, b...)
	return len(b), nil
}

// TestStartServer tests the TCP server.
func TestStartServer(t *testing.T) {
	go startServer() // Start the server in a goroutine

	time.Sleep(1 * time.Second) // Wait a moment for the server to start

	// Try to connect to the server
	conn, err := net.Dial("tcp", serverAddress)
	if err != nil {
		t.Fatalf("Failed to connect to server: %v", err)
	}
	defer conn.Close()

	// If connection is successful, the test passes
	t.Log("Successfully connected to the server")
}

/*
	 func TestHandleConnection(t *testing.T) {
		// Create a pipe to simulate a connection
		serverConn, clientConn := net.Pipe()
		defer serverConn.Close()
		defer clientConn.Close()

		// Run the handleConnection function in a goroutine
		go handleConnection(serverConn)

		// Simulate client sending data
		testMessage := []byte("Hello, Server!")
		_, err := clientConn.Write(testMessage)
		if err != nil {
			t.Fatalf("Failed to write to client connection: %v", err)
		}

		// You can add logic here to verify the server's response
		// For example, if handleConnection sends a response back,
		// you would read from the clientConn to check the response.
		// Here's a basic example of reading a response:

		buf := make([]byte, 1024)
		n, err := clientConn.Read(buf)
		if err != nil && err != io.EOF {
			t.Fatalf("Failed to read from client connection: %v", err)
		}

		// Here, we could add assertions based on what we expect
		// to receive from handleConnection after sending the message.

		t.Logf("Received from server: %s", string(buf[:n]))
	}
*/
func TestCreateUser(t *testing.T) {
	// make the fake connection doodad
	serverConn, clientConn := net.Pipe()
	defer serverConn.Close()
	defer clientConn.Close()

	packet := Packet{Username: "shelwig", Address: clientConn.LocalAddr()}

	result := createUser(packet)

	expected := User{Username: "shelwig", Address: clientConn.LocalAddr()}

	if result != expected {
		t.Errorf("Created user does not equal expected user")
	}

}

func TestSerialization(t *testing.T) {
	// pass a packet in, serialize it, deserialize, make it equal
	packet := Packet{Username: "shelwig"}

	data := serializePacket(packet)
	newPacket := deserializePacket(data)

	if packet != newPacket {
		t.Errorf("Serialization failed - new packet does not equal old")
	}

}

func TestInitialConnection(t *testing.T) {
	// connect
	// send the packet of initial connection
	serverConn, clientConn := net.Pipe()
	defer serverConn.Close()
	defer clientConn.Close()

	startServer()

	packet := Packet{Username: "shelwig", Goal: "initial_connection", Address: clientConn.LocalAddr()}
	user := User{Username: "shelwig1", Address: clientConn.LocalAddr()}

	data := serializePacket(packet)

	// Connect to the server and send that packet

	conn, err := net.Dial("tcp", serverAddress)
	if err != nil {
		t.Fatalf("Failed to connect to server: %v", err)
	}
	defer conn.Close()

	conn.Write()

	expected := []User{user}
	result := active_users

	if len(expected) != len(result) {
		t.Errorf("Lists not same length")
	}
	for i := range expected {
		if expected[i].Username != result[i].Username || expected[i].Address != result[i].Address {
			t.Errorf("Users not the same")

		}
	}

}
