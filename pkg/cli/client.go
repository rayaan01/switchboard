package cli

import (
	"bufio"
	"fmt"
	"io"
	"net"
	"os"
	"time"
)

type Client struct {
	address   string
	conn      net.Conn
	accessKey string
}

func CreateClient(host *string, port *string) (*Client, error) {
	address := fmt.Sprintf("%s:%s", *host, *port)
	conn, err := net.Dial("tcp", address)
	accessKey := "EMPTY"
	if err != nil {
		return nil, err
	}
	clientInstance := Client{address, conn, accessKey}
	return &clientInstance, nil
}

func (c *Client) HandleConnection() {
	reader := bufio.NewScanner(os.Stdin)
	for reader.Scan() {
		input := reader.Text()
		if len(input) == 0 {
			displayWrapper("")
			continue
		}
		start := time.Now()
		response, err := c.parser(input)
		duration := time.Since(start)
		if err != nil {
			if err == io.EOF {
				displayWrapper("The connection has been closed. Thank you!\n")
				return
			}
			displayWrapper(err.Error())
			continue
		}
		formattedResponse := []byte(fmt.Sprintf(" (%s)\n", duration))
		wrappedResponse := string(append(response, formattedResponse...))
		displayWrapper(wrappedResponse)
		displayWrapper("")
	}
}
