package cli

import (
	"bufio"
	"fmt"
	"io"
	"net"
	"os"
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
		response, err := c.parser(input)
		if err != nil {
			if err == io.EOF {
				displayWrapper("The connection has been closed. Thank you!\n")
				return
			}
			displayWrapper(err.Error())
			continue
		}
		wrappedResponse := string(append(response, []byte("\n")...))
		displayWrapper(wrappedResponse)
		displayWrapper("")
	}
}
