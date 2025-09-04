package main

import (
	"bufio"
	"context"
	"net"
	"sync"
	"testing"
)

func TestHandleConnection(t *testing.T) {
	// Create a mock connection using net.Pipe
	serverConn, clientConn := net.Pipe()

	var wg sync.WaitGroup
	ctx, cancel := context.WithCancel(context.Background())

	wg.Add(1)
	go handleConnection(serverConn, &wg, ctx)

	// Client sends a message
	clientConn.Write([]byte("hello\n"))

	// Read the echo
	scanner := bufio.NewScanner(clientConn)
	if !scanner.Scan() {
		t.Fatal("Expected echo response")
	}
	response := scanner.Text()
	if response != "Echo: hello" {
		t.Errorf("Expected 'Echo: hello', got '%s'", response)
	}

	// Cancel context to shutdown
	cancel()

	// Wait for handler to finish
	wg.Wait()

	// Close client conn
	clientConn.Close()
}

func TestServerAcceptsConnections(t *testing.T) {
	// This test is harder to implement without mocking or running in goroutine
	// For now, skip full server test
	t.Skip("Full server test requires more setup")
}
