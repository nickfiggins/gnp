package main

import "testing"

func TestZeroWindow(t *testing.T) {
	/*

	More info: https://accedian.com/blog/tcp-receive-window-everything-need-know/

	Def'n: when a TCP buffer is full and can not receive more data, due to a stuck processor or busy with another task

	buf := make([]byte, 1024)
	for {
		n, err := conn.Read(buf) reading frees up buffer space
		if err != nil {
			return err
		}
		handle(buf[:n]) // BLOCKS!
	}
	*/
}