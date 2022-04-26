package echo

import (
	"context"
	"fmt"
	"log"
	"net"
	"os"
)

func echoServerUDP(ctx context.Context, addr string) (net.Addr, error) {
	s, err := net.ListenPacket("udp", addr)
	flags := log.Lshortfile | log.Ldate | log.Ltime | log.Lmicroseconds
	logger := log.New(os.Stderr, "", flags)
	
	if err != nil {
		return nil, fmt.Errorf("binding to udp %s: %w", addr, err)
	}

	go func() {
		go func() {
			<-ctx.Done()
			_ = s.Close()
		}()

		buf := make([]byte, 1024)
		for {
			n, clientAddr, err := s.ReadFrom(buf) // client to server
			logger.Printf("Read from client: %v %v %q", n, clientAddr, err)
			if err != nil {
				return
			}

			_, err = s.WriteTo(buf[:n], clientAddr) // server to client
			logger.Printf("Wrote to client: %v %v %q", n, clientAddr, err)
			if err != nil {
				return
			}
		}
	}()

	return s.LocalAddr(), nil
}
