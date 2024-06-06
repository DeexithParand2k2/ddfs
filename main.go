package main

import (
	"log"

	"github.com/DeexithParand2k2/ddfs/p2p"
)

// check if outbound peer details are printed after conn established via
// ... powershell telnet localhost:5000
func main() {

	var listenAddr string = ":5000"

	tcpServer := p2p.NewTCPTransport(listenAddr)
	err := tcpServer.ListenAndAccept()
	if err != nil {
		log.Fatal("Error listening to the server at : ", listenAddr)
	}

	// used to create an infinite loop that waits for communication on multiple channels without consuming CPU resources unnecessarily.
	select {}

}
