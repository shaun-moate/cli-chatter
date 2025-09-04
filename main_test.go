package main

import (
	"bufio"
	"context"
	"fmt"
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

func TestRunClientConnectionFailure(t *testing.T) {
	// Test connection failure
	err := runClient("invalid:address")
	if err == nil {
		t.Error("Expected error for invalid address")
	}
}

func TestRunClientSuccessfulConnection(t *testing.T) {
	// Start a test server
	ln, err := net.Listen("tcp", "localhost:0")
	if err != nil {
		t.Fatal(err)
	}
	defer ln.Close()

	addr := ln.Addr().String()

	// Run client in goroutine
	go func() {
		err := runClient(addr)
		if err != nil {
			t.Errorf("Client error: %v", err)
		}
	}()

	// Accept connection
	conn, err := ln.Accept()
	if err != nil {
		t.Fatal(err)
	}
	defer conn.Close()

	// Send a message from server
	fmt.Fprintf(conn, "Hello from server\n")

	// For simplicity, just check that connection was accepted
	// In real test, could check input forwarding, but requires mocking stdin
}
