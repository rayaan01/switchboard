package cli

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"strings"

	"switchboard/pkg/cli/utils"
	"switchboard/pkg/common"
	"switchboard/pkg/types"
)

func (c *Client) parser(input string) ([]byte, error) {
	usageMessage := common.GetUsageMessage()
	args := strings.Fields(input)
	serializedCmd := strings.Join(args, " ")
	cmdType := args[0]

	switch cmdType {
	case "set":
		if len(args) != 3 {
			return usageMessage, nil
		}

	case "get", "del":
		if len(args) != 2 {
			return usageMessage, nil
		}

	case "create-access-key":
		if len(args) != 2 {
			return usageMessage, nil
		}
		if args[1] != "HashTable" && args[1] != "AVLTree" {
			return usageMessage, nil
		}

	case "use":
		if len(args) != 2 {
			return usageMessage, nil
		}

		accessKey := args[1]
		response, err := utils.UpdateKeys(accessKey)
		if err != nil {
			return nil, err
		}
		c.accessKey = accessKey
		return response, nil

	case "visualize":
		if len(args) != 1 {
			return usageMessage, nil
		}

	case "exit":
		return nil, io.EOF

	default:
		return usageMessage, nil
	}

	if c.accessKey == "EMPTY" {
		accessKey, err := utils.GetAccessKey()
		if err != nil {
			return nil, err
		}
		c.accessKey = accessKey
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
