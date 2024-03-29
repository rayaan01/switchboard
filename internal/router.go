package internal

import (
	"io"
	"strings"
)

func router(args []string) ([]byte, error) {
	errorMessage := []byte("Available commands:\n1. set [key] [value]\n2. get [key]\n3. del [key]\n4. exit")
	if len(args) == 0 {
		return errorMessage, nil
	}
	cmd := strings.ToLower(args[0])
	switch cmd {
	case "set":
		if len(args) != 3 {
			return errorMessage, nil
		}
		key := args[1]
		val := args[2]
		response, err := handleSet(key, val)
		if err != nil {
			return nil, err
		}
		return append(response, []byte("\n")...), nil
	case "get":
		if len(args) != 2 {
			return errorMessage, nil
		}
		key := args[1]
		response, err := handleGet(key)
		if err != nil {
			return nil, err
		}
		return append(response, []byte("\n")...), nil
	case "del":
		if len(args) != 2 {
			return errorMessage, nil
		}
		key := args[1]
		response, err := handleDel(key)
		if err != nil {
			return nil, err
		}
		return append(response, []byte("\n")...), nil
	case "exit":
		return nil, io.EOF
	default:
		return errorMessage, nil
	}
}
