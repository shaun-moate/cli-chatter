package main

import (
	"bufio"
	"context"
	"flag"
	"fmt"
	"net"
	"os"
	"os/signal"
	"sync"
	"syscall"
)

func main() {
	serverFlag := flag.Bool("server", false, "Run in server mode")
	clientFlag := flag.String("client", "", "Run in client mode, specify server address (e.g., localhost:8080)")
	flag.Parse()

	if *serverFlag && *clientFlag != "" {
		fmt.Fprintf(os.Stderr, "Error: Cannot specify both --server and --client\n")
		flag.Usage()
		os.Exit(1)
	}

	if !*serverFlag && *clientFlag == "" {
		fmt.Fprintf(os.Stderr, "Error: Must specify either --server or --client\n")
		flag.Usage()
		os.Exit(1)
	}

	if *serverFlag {
		runServer()
	} else {
		if err := runClient(*clientFlag); err != nil {
			os.Exit(1)
		}
	}
}

func runServer() {
	fmt.Fprintf(os.Stderr, "Starting server mode, listening on :8080\n")
	ln, err := net.Listen("tcp", ":8080")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error starting server: %v\n", err)
		os.Exit(1)
	}
	defer ln.Close()

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	var wg sync.WaitGroup

	// Handle shutdown signals
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)
	go func() {
		<-sigChan
		fmt.Fprintf(os.Stderr, "Shutting down server...\n")
		cancel()
		ln.Close()
	}()

	for {
		select {
		case <-ctx.Done():
			fmt.Fprintf(os.Stderr, "Stopping to accept new connections\n")
			goto waitForConnections
		default:
		}

		conn, err := ln.Accept()
		if err != nil {
			select {
			case <-ctx.Done():
				fmt.Fprintf(os.Stderr, "Listener closed\n")
			default:
				fmt.Fprintf(os.Stderr, "Error accepting connection: %v\n", err)
			}
			continue
		}
		wg.Add(1)
		go handleConnection(conn, &wg, ctx)
	}

waitForConnections:
	wg.Wait()
	fmt.Fprintf(os.Stderr, "All connections closed. Server shutdown complete.\n")
}

func handleConnection(conn net.Conn, wg *sync.WaitGroup, ctx context.Context) {
	defer wg.Done()
	defer conn.Close()

	done := make(chan struct{})
	go func() {
		defer close(done)
		scanner := bufio.NewScanner(conn)
		for scanner.Scan() {
			msg := scanner.Text()
			fmt.Fprintf(os.Stderr, "Received: %s\n", msg)
			// Echo back for simplicity
			fmt.Fprintf(conn, "Echo: %s\n", msg)
		}
		if err := scanner.Err(); err != nil {
			fmt.Fprintf(os.Stderr, "Error reading from connection: %v\n", err)
		}
	}()

	select {
	case <-ctx.Done():
		fmt.Fprintf(os.Stderr, "Shutting down connection\n")
		conn.Close() // This will cause scanner.Scan() to return
	case <-done:
		// Connection closed naturally
	}
}

func runClient(serverAddr string) error {
	fmt.Fprintf(os.Stderr, "Starting client mode, connecting to %s\n", serverAddr)
	conn, err := net.Dial("tcp", serverAddr)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error connecting to server: %v\n", err)
		return err
	}
	defer conn.Close()

	// Goroutine to read from server
	go func() {
		scanner := bufio.NewScanner(conn)
		for scanner.Scan() {
			fmt.Println(scanner.Text())
		}
		if err := scanner.Err(); err != nil {
			fmt.Fprintf(os.Stderr, "Error reading from server: %v\n", err)
		}
	}()

	// Read from stdin and send to server
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		input := scanner.Text()
		fmt.Fprintf(conn, "%s\n", input)
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintf(os.Stderr, "Error reading input: %v\n", err)
		return err
	}
	return nil
}
