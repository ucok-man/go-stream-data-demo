package main

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"io"
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

	buff := new(bytes.Buffer)

	// Because we don't know how many bytes we're going to receive, we're going to use CopyN to read the bytes.
	// But we need to know how many bytes to read, we're going to read the size from the connection first.
	// We're going to use binary.Read to read the size from the connection.

	var size int64
	if err := binary.Read(conn, binary.LittleEndian, &size); err != nil {
		slog.Error("Error reading size bytes", slog.Any("error", err))
	}

	_, err := io.CopyN(buff, conn, size)
	if err != nil && err.Error() != "EOF" {
		slog.Error("Error reading into buffer", slog.Any("error", err))
	}

	fmt.Println("---------------------------------")
	fmt.Println("data chunk:")
	fmt.Println(buff.Bytes())

	fmt.Println("---------------------------------")
	slog.Info("Finished reading data", slog.Any("size_bytes", size))
}
