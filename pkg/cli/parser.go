package cli

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"os"
	"strings"
)

type Request struct {
	Key string `json:"key"`
	Cmd string `json:"cmd"`
}

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
		existingKeys, err := os.ReadFile("keys.json")

		if err != nil {
			keyStructure := map[string]bool{accessKey: true}
			marshalledKey, err := json.Marshal(keyStructure)
			if err != nil {
				return nil, err
			}
			os.WriteFile("keys.json", marshalledKey, 0644)
			return []byte("OK"), nil
		}

		var unmarshalledKeys map[string]bool
		err = json.Unmarshal(existingKeys, &unmarshalledKeys)
		if err != nil {
			return nil, err
		}
		for k := range unmarshalledKeys {
			unmarshalledKeys[k] = false
		}
		unmarshalledKeys[accessKey] = true
		marshalledKeys, err := json.Marshal(unmarshalledKeys)
		if err != nil {
			return nil, err
		}
		os.WriteFile("keys.json", marshalledKeys, 0644)
		return []byte("OK"), nil

	case "exit":
		return nil, io.EOF
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
