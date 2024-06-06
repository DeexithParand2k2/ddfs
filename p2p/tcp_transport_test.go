package p2p

import (
	"net"
	"testing"
)

func TestTCPTransport(t *testing.T) {

	var listenAddr string = ":8000"

	// tcpServerNode := NewTCPTransport("localhost:8000")

	// bind server to port 8000 and allow connection from all Network Interfaces
	// ... not only loopback local addr, but also to be able to connect to devices on same network
	// ... allow server to listen to connection from LAN AND WAN
	tcpServerNode := NewTCPTransport(listenAddr)

	have := tcpServerNode
	expected := make(map[net.Addr]Peer)

	// expecting same listening addr and 0 peers on conn initiation
	if tcpServerNode.listenAddress == listenAddr && len(have.peers) == len(expected) {
		t.Logf("test passed 0 peers on connection initiation")
	} else {
		t.Fatalf("Error initating a tcp server")
	}

	// listen and accept
	err := tcpServerNode.ListenAndAccept()
	if err != nil {
		t.Fatalf("Error initating a tcp server")
	} else {
		t.Logf("test passed with server listening at port : %s", tcpServerNode.listenAddress)
	}

}
