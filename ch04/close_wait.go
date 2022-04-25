package main

import (
	"testing"
)



func TestCloseWait(t *testing.T) {
	/*

	Leaves TCP connection in CLOSE_WAIT state


	for {
		conn, err := listener.Accept()
		if err != nil {
			return err
		}
		go func(c net.Conn) { // we never call c.Close() before returning!
			buf := make([]byte, 1024)
			for {
				n, err := c.Read(buf)
				if err != nil {
					return
				}
				handle(buf[:n])
			}
		}(conn)
	}
	*/
}