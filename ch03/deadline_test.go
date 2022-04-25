package ch03

import (
	"io"
	"net"
	"testing"
	"time"
)

func TestDeadline(t *testing.T) {
	sync := make(chan struct{})

	listener, err := net.Listen("tcp", "127.0.0.1:")
	if err != nil {
		t.Fatal(err)
	}

	go func() {
		conn, err := listener.Accept()
		if err != nil {
			t.Log(err)
			return
		}
		defer func() {
			conn.Close()
			close(sync) // read from sync shouldn't block due to early return
		}()

		err = conn.SetDeadline(time.Now().Add(5 * time.Second)) // 5 second deadline to read (and/or write)
		if err != nil {
			t.Error(err)
			return
		}

		buf := make([]byte, 1)
		t.Log("Beginning read...")
		_, err = conn.Read(buf) // blocked until remote node sends data
		t.Log("Finished reading!") // 5 seconds later
		nErr, ok := err.(net.Error) // we get a timeout since no data was sent after 5 seconds
		if !ok || !nErr.Timeout() {
			t.Errorf("expected timeout error; actual: %v", err)
		}

		sync <- struct{}{}

		err = conn.SetDeadline(time.Now().Add(5 * time.Second))
		if err != nil {
			t.Error(err)
			return
		}

		_, err = conn.Read(buf) // wait 5 seconds again
		if err != nil {
			t.Error(err)
		}

		if string(buf) != "1" {
			t.Error("buf not equal to written value")
		}
	}()

	conn, err := net.Dial("tcp", listener.Addr().String())
	if err != nil {
		t.Fatal(err)
	}
	defer conn.Close()

	<-sync // triggers the write beginning on line 44
	_, err = conn.Write([]byte("1"))
	if err != nil {
		t.Fatal(err)
	}

	buf := make([]byte, 1)
	_, err = conn.Read(buf) // blocked until remote node sends data
	if err != io.EOF {
		t.Errorf("expected server termination; actual: %v", err)
	}
}
