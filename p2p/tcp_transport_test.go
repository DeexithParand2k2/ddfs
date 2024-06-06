package p2p

import (
	"fmt"
	"net"
	"testing"
)

func TestTCPTransport(t *testing.T) {
	// tcpServerNode := NewTCPTransport("localhost:8000")

	// bind server to port 8000 and allow connection from all Network Interfaces
	// ... not only loopback local addr, but also to be able to connect to devices on same network
	// ... allow server to listen to connection from LAN AND WAN
	tcpServerNode := NewTCPTransport(":8000")

	have := tcpServerNode.peers
	expected := make(map[net.Addr]Peer)

	fmt.Println("Initiated tcp Server : ", tcpServerNode)

	if len(have) == len(expected) && len(expected) == 0 {
		t.Logf("test passed 0 peers connection init")
		return
	}

	t.Fatalf("Error initating a tcp server")
}
