package main

import (
	"bufio"
	"flag"
	"fmt"
	"net"
	"os"
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
		runClient(*clientFlag)
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

	for {
		conn, err := ln.Accept()
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error accepting connection: %v\n", err)
			continue
		}
		go handleConnection(conn)
	}
}

func handleConnection(conn net.Conn) {
	defer conn.Close()
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
}

func runClient(serverAddr string) {
	fmt.Fprintf(os.Stderr, "Starting client mode, connecting to %s\n", serverAddr)
	conn, err := net.Dial("tcp", serverAddr)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error connecting to server: %v\n", err)
		os.Exit(1)
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
	}
}
