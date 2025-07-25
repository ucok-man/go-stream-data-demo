package main

import (
	"bytes"
	"crypto/rand"
	"encoding/binary"
	"io"
	"log/slog"
	"net"
	"os"
)

func main() {
	conn, err := net.Dial("tcp", "localhost:4000")
	if err != nil {
		slog.Error("Error connecting to server", slog.Any("error", err))
		os.Exit(1)
	}
	defer conn.Close()

	file := make([]byte, 10000)
	_, err = io.ReadFull(rand.Reader, file)
	if err != nil {
		slog.Error("Error mocking file", slog.Any("error", err))
		os.Exit(1)
	}

	// First we tell to server the full size of the file will be
	if err := binary.Write(conn, binary.LittleEndian, int64(len(file))); err != nil {
		slog.Error("Error writing file to server", slog.Any("error", err))
		os.Exit(1)
	}
	// Copy will copy to target chunk by chunk until find EOF
	n, err := io.Copy(conn, bytes.NewReader(file))
	if err != nil {
		slog.Error("Error writing file to server", slog.Any("error", err))
		os.Exit(1)
	}

	slog.Info("Success writing file!!!", slog.Any("size_bytes", n))
}
