package db

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"io"
	"net"
	"os"
	"strings"

	"switchboard/pkg/types"
)

func Handler(conn net.Conn, s *Server) {
	defer conn.Close()
	filePath := "metrics.csv"
	file, err := os.OpenFile(filePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		fmt.Printf("Error opening %s file: %s", filePath, err)
	}
	metricsWriter := csv.NewWriter(file)
	defer file.Close()

	for {
		buffer := make([]byte, 0, 1024)
		bytesRead := 0
		clientAddress := conn.RemoteAddr().String()
		err := readCommand(&buffer, &bytesRead, conn)
		if err != nil {
			if err == io.EOF {
				return
			}
			fmt.Printf("Could not read from connection %s : %s \n", clientAddress, err)
			conn.Write([]byte("Something went wrong!"))
			continue
		}
		var request types.Request
		json.Unmarshal(buffer[:bytesRead], &request)
		args := strings.Fields(request.Cmd)
		response, err := router(request.Key, args, metricsWriter)
		if err != nil {
			if err == io.EOF {
				return
			}
			fmt.Printf("Router error on %s : %s \n", clientAddress, err)
			conn.Write([]byte("Something went wrong!"))
			continue
		}
		conn.Write(response)
	}
}

func readCommand(buffer *[]byte, bytesRead *int, conn net.Conn) error {
	for {
		chunk := make([]byte, 128)
		n, err := conn.Read(chunk)
		if err != nil {
			return err
		}
		*buffer = append(*buffer, chunk[:n]...)
		*bytesRead += n
		if n < cap(chunk) {
			return nil
		}
	}
}
