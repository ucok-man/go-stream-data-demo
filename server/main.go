package main

import (
	"fmt"
	"log/slog"
	"net"
	"os"
)

func main() {
	listener, err := net.Listen("tcp", ":4000")
	if err != nil {
		slog.Error("Error listening tcp", slog.Any("error", err))
		os.Exit(1)
	}
	defer listener.Close()
	slog.Info("Server started", slog.Any("port", 4000))

	for {
		conn, err := listener.Accept()
		if err != nil {
			slog.Error("Error accepting connection", slog.Any("error", err))
			continue
		}
		go handleConnection(conn)
	}
}

func handleConnection(conn net.Conn) {
	defer conn.Close()
	slog.Info("Accepted connection", slog.Any("remote_addr", conn.RemoteAddr()))

	buff := make([]byte, 2048)
	totalBytes := 0

	for {
		n, err := conn.Read(buff)
		if err != nil {
			if err.Error() == "EOF" {
				break
			}
			slog.Error("Error reading into buffer", slog.Any("error", err))
		}

		fmt.Println("---------------------------------")
		fmt.Println("data chunk:")
		fmt.Println(buff[:n])

		totalBytes += n
	}

	fmt.Println("---------------------------------")
	slog.Info("Finished reading data", slog.Any("size_bytes", totalBytes))
}
