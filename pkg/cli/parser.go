package cli

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"net"
	"strings"
)

func Parser(input string) (string, error) {
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
		return "", io.EOF
	case "default":
		return getUsageMessage(), nil
	}

	response, err := sendCommand(serializedCmd)
	if err != nil {
		return "", err
	}
	return response, nil
}

func sendCommand(cmd string) (string, error) {
	conn, err := net.Dial("tcp", "localhost:8080")
	if err != nil {
		msg := fmt.Sprintf("%s %s", "Could not connect to server", err)
		return "", errors.New(msg)
	}
	_, err = fmt.Fprint(conn, cmd)
	if err != nil {
		msg := fmt.Sprintf("%s %s", "Could not write to connection", err)
		return "", errors.New(msg)
	}
	response, err := bufio.NewReader(conn).ReadString('\n')
	if err != nil {
		if err != io.EOF {
			msg := fmt.Sprintf("%s %s", "Connection interrupted by server", err)
			return "", errors.New(msg)
		}
		return "", io.EOF
	}
	return response, nil
}
