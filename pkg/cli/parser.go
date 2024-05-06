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

	default:
		return getUsageMessage(), nil
	}

	request := types.Request{Key: c.accessKey, Cmd: serializedCmd}
	marshalledRequest, err := json.Marshal(request)
	if err != nil {
		return nil, err
	}
	response, err := c.sendRequest(marshalledRequest)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (c *Client) sendRequest(request []byte) ([]byte, error) {
	_, err := c.conn.Write(request)
	if err != nil {
		msg := fmt.Sprintf("%s %s", "Could not write to connection", err)
		return nil, errors.New(msg)
	}
	response := make([]byte, 1024)
	bytesRead, err := c.conn.Read(response)
	if err != nil {
		return nil, err
	}
	return response[:bytesRead], nil
}
