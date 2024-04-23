package cli

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"strings"

	"github.com/google/uuid"
)

func (c *Client) parser(input string) ([]byte, error) {
	args := strings.Fields(input)
	serializedCmd := strings.Join(args, " ")
	if len(args) == 0 {
		return getUsageMessage(), nil
	}

	cmd := args[0]
	switch cmd {
	case "set":
		if len(args) != 3 {
			return getUsageMessage(), nil
		}
	case "get", "del":
		if len(args) != 2 {
			return getUsageMessage(), nil
		}
	case "exit":
		return nil, io.EOF
	case "CreateAccessKey":
		accessKey := []byte(uuid.NewString())
		return accessKey, nil
	case "default":
		return getUsageMessage(), nil
	}

	response, err := c.sendCommand(serializedCmd)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (c *Client) sendCommand(cmd string) ([]byte, error) {
	_, err := fmt.Fprint(c.conn, cmd)
	if err != nil {
		msg := fmt.Sprintf("%s %s", "Could not write to connection", err)
		return nil, errors.New(msg)
	}
	response := make([]byte, 0, 128)
	bytesRead := 0
	err = c.readResponse(&response, &bytesRead)
	if err != nil {
		fmt.Printf("Could not read response: %s", err)
		return []byte("Something went wrong!"), nil
	}
	if err != nil {
		if err != io.EOF {
			msg := fmt.Sprintf("%s %s", "Connection interrupted by server", err)
			return nil, errors.New(msg)
		}
		return nil, io.EOF
	}
	return response[:bytesRead], nil
}

func (c *Client) readResponse(responseBuffer *[]byte, bytesRead *int) error {
	for {
		chunk := make([]byte, 32)
		n, err := bufio.NewReader(c.conn).Read(chunk)
		if err != nil {
			return err
		}
		*responseBuffer = append(*responseBuffer, chunk[:n]...)
		*bytesRead += n
		if n < cap(chunk) {
			return nil
		}
	}
}
