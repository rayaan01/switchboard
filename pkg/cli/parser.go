package cli

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"strings"

	"switchboard/pkg/cli/handlers"
	"switchboard/pkg/types"
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

	case "create-access-key":
		if len(args) != 2 {
			return getUsageMessage(), nil
		}
		if args[1] != "HashTable" && args[1] != "AVLTree" {
			return getUsageMessage(), nil
		}

	case "use":
		if len(args) != 2 {
			return getUsageMessage(), nil
		}

		accessKey := args[1]
		response, err := handlers.HandleUse(accessKey)
		if err != nil {
			return nil, err
		}
		c.accessKey = accessKey
		return response, nil

	case "exit":
		return nil, io.EOF

	case "default":
		return getUsageMessage(), nil
	}

	request := types.Request{Key: c.accessKey, Cmd: serializedCmd}
	marshalledRequest, err := json.Marshal(request)
	if err != nil {
		return nil, err
	}
	serializedRequest := string(marshalledRequest)
	response, err := c.sendCommand(serializedRequest)
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

	response := make([]byte, 0, 1024)
	bytesRead := 0
	err = c.readResponse(&response, &bytesRead)
	if err != nil {
		return nil, err
	}
	return response[:bytesRead], nil
}

func (c *Client) readResponse(responseBuffer *[]byte, bytesRead *int) error {
	for {
		chunk := make([]byte, 128)
		n, err := c.conn.Read(chunk)
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
