package p2p

import (
	"errors"
	"log"
	"net"
	"sync"
)

type TCPPeer struct {
	conn net.Conn

	// dial a connection => true
	// accept a connection => false
	outbound bool
}

// constructor to TCPPeer
func NewTCPPeer(conn net.Conn, outbound bool) *TCPPeer {
	return &TCPPeer{
		conn:     conn,
		outbound: true,
	}
}

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

// Method Signature : Initiate server that listens and Accepts connectin at that network address
func (tcpServer *TCPTransport) ListenAndAccept() error {
	var err error
	tcpServer.listener, err = net.Listen("tcp", tcpServer.listenAddress)
	if err != nil {
		log.Println("Error listening to server at", tcpServer.listener, " ", err)
		return err
	}

	// trigger a go routine : initiate server to accept connection
	go tcpServer.StartAcceptLoop()

	log.Printf("TCP transport listening on port: %s\n", tcpServer.listenAddress)

	// Return nil to indicate that there were no errors during setup
	return nil
}

// Method Signature that accepts infinitely accepts tcp connection until listener is closed
func (tcpServer *TCPTransport) StartAcceptLoop() {

	for {
		tcpConn, err := tcpServer.listener.Accept()
		if errors.Is(err, net.ErrClosed) { // ErrClosed is the error returned by an I/O call on a network
			return
		}

		if err != nil {
			log.Printf("Error Accepting tcp connection")
		}

		go tcpServer.handleTCPConnection(tcpConn)
	}

}

func (tcpServer *TCPTransport) handleTCPConnection(conn net.Conn) {

	peer := NewTCPPeer(conn, true)

	log.Println("Listener accepted tcp connection : ", peer)
}
