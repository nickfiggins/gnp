package ch03

import (
	"context"
	"net"
	"syscall"
	"testing"
	"time"
)

func TestDialContext(t *testing.T) {
	dl := time.Now().Add(5 * time.Second)
	ctx, cancel := context.WithDeadline(context.Background(), dl)
	defer cancel()

	var d net.Dialer // DialContext is a method on a Dialer
	d.Control = func(_, _ string, _ syscall.RawConn) error {
		// Sleep long enough to reach the context's deadline.
		time.Sleep(5*time.Second + time.Millisecond)
		return nil
	}
	conn, err := d.DialContext(ctx, "tcp", "10.0.0.0:80")
	if err == nil {
		conn.Close()
		t.Fatal("connection did not time out")
	}
	nErr, ok := err.(net.Error)
	if !ok {
		t.Error(err)
	} else {
		if !nErr.Timeout() {
			t.Errorf("error is not a timeout: %v", err)
		}
	}
	if ctx.Err() != context.DeadlineExceeded {
		t.Errorf("expected deadline exceeded; actual: %v", ctx.Err())
	}
}

func TestDialContext2(t *testing.T) {
	// create a deadline of five seconds into the future
	dl := time.Now().Add(5 * time.Second)

	// create a context with a 5 seconds deadline into the future
	// and get the cancel function
	ctx, cancel := context.WithDeadline(context.Background(), dl)

	// it's a good practice to defer the cancel function to make sure the
	// context is garbage collected as soon as possible.
	defer cancel()

	var d net.Dialer

	// overrides the Control function of the dialer
	// delays the connection long enough to ensure we exceed the
	// context deadline (5.001s)
	d.Control = func(_, _ string, _ syscall.RawConn) error {
		// sleep long enough to reach the context deadline
		time.Sleep(5*time.Second + time.Millisecond)
		return nil
	}

	// pass the context (ctx) to the DialContext function
	// of the dialer
	conn, err := d.DialContext(ctx, "tcp", "10.0.0.0:80")
	if err != nil {
		if conn != nil {
			conn.Close()
		}
		t.Fatal("connection did not time out")
	}
}
