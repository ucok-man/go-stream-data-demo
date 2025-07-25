package main

import (
	"crypto/rand"
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

	file := make([]byte, 4000)
	_, err = io.ReadFull(rand.Reader, file)
	if err != nil {
		slog.Error("Error mocking file", slog.Any("error", err))
		os.Exit(1)
	}

	n, err := conn.Write(file)
	if err != nil {
		slog.Error("Error writing file to server", slog.Any("error", err))
		os.Exit(1)
	}

	slog.Info("Success writing file!!!", slog.Any("size_bytes", n))
}
