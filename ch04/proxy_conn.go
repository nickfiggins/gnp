package main

import (
	"io"
	"net"
)


// net.Conn implements both the io.Writer and io.Reader interface, which allows us to easily us io.Copy to proxy requests
func proxyConn(source, destination string) error {
	connSource, err := net.Dial("tcp", source)
	if err != nil {
		return err
	}
	defer connSource.Close()

	connDestination, err := net.Dial("tcp", destination)
	if err != nil {
		return err
	}
	defer connDestination.Close()

	// connSource <- connDestination (replies)
	go func() { _, _ = io.Copy(connSource, connDestination) }()

	// connDestination <- connSource
	_, err = io.Copy(connDestination, connSource)

	return err
}

var _ = proxyConn
