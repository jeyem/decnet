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

func (c *Connection) newContext() *Context {
	ctx := &Context{}

	ctx.request = bufio.NewBuffer(ctx.requestBuffer)
	ctx.response = bufio.NewBuffer(ctx.responseBuffer)

	return ctx
}

func (ctx *Context) Replay(message string) error {
	_, err := ctx.response.Write([]byte(message))
	return err
}
