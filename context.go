package decnet

import (
	"net"

	"gopkg.in/bufio.v1"
)

type Context struct {
	requestBuffer  []byte
	responseBuffer []byte
	request        *bufio.Buffer
	response       *bufio.Buffer
	connection     *Connection
	tcpConnection  net.Conn
}

func newContext() *Context {
	c := &Context{
		requestBuffer:  make([]byte, 2048),
		responseBuffer: make([]byte, 2048),
	}

	c.request = bufio.NewBuffer(c.requestBuffer)
	c.response = bufio.NewBuffer(c.responseBuffer)

	return c
}

func (c *Context) Replay(message string) error {
	_, err := c.response.Write([]byte(message))
	return err
}
