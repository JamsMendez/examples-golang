package tcp

import (
	"net"
	"testing"
	"time"
)

func TestServerTCP(t *testing.T) {
	s, err := NewServerTCP(":8080")
	if err != nil {
		t.Fatal(err)
	}

	s.Start()

	var conn net.Conn
	conn, err = net.Dial("tcp", "localhost:8080")
	if err != nil {
		t.Fatal(err)
	}

	defer conn.Close()

	expected := "Welcome to TCP Server!\n"
	actual := make([]byte, len(expected))
	if _, err = conn.Read(actual); err != nil {
		t.Fatal(err)
	}

	if string(actual) != expected {
		t.Errorf("expected %q, but got %q", expected, actual)
	}

	<-time.After(6 * time.Second)

	expected = "Goodbye!\n"
	actual = make([]byte, len(actual))
	if _, err = conn.Read(actual); err != nil {
		t.Fatal(err)
	}

	if string(actual) != expected {
		t.Errorf("expected %q, but got %q", expected, actual)
	}

	s.Stop()
}
