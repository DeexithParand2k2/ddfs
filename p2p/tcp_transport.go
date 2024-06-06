package p2p

import (
	"net"
	"sync"
)

// TCPTransport manages TCP network connections and handles the communication
// between peers in a peer-to-peer network.
type TCPTransport struct {
	listenAddress string       // The address where the TCP server listens for incoming connections.
	listener      net.Listener // The network listener that accepts incoming connections.

	mu    sync.RWMutex      // A read/write mutex to ensure thread-safe access to the peers map.
	peers map[net.Addr]Peer // A map to keep track of connected peers, indexed by their network address.
}

// Constructor to the TCPTransport Struct
// NewTCPTransport creates and initializes a new TCPTransport instance with the specified
// listening address. It sets up the address and initializes the peers map.
func NewTCPTransport(listenAddr string) *TCPTransport {
	return &TCPTransport{
		listenAddress: listenAddr,              // Set the address where the server will listen for connections.
		peers:         make(map[net.Addr]Peer), // Initialize the map to store connected peers.
	}
}
