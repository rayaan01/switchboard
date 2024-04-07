package cli

import (
	"bufio"
	"fmt"
	"io"
	"net"
	"os"
)

type Client struct {
	address string
	conn    net.Conn
}

func CreateClient(host *string, port *string) (*Client, error) {
	address := fmt.Sprintf("%s:%s", *host, *port)
	conn, err := net.Dial("tcp", address)
	if err != nil {
		return nil, err
	}
	clientInstance := Client{address, conn}
	return &clientInstance, nil
}

func (c *Client) HandleConnection() {
	reader := bufio.NewScanner(os.Stdin)
	for reader.Scan() {
		input := reader.Text()
		response, err := c.parser(input)
		if err != nil {
			if err == io.EOF {
				fmt.Println("The connection has been closed. Thank you!")
				return
			}
			fmt.Println(err)
			continue
		}
		displayWrapper(response)
		displayWrapper("")
	}
}
