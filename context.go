package decnet

import (
	"io"
	"net"
	"sync"

	"gopkg.in/bufio.v1"
)

type Context struct {
	requestBuffer  []byte
	responseBuffer []byte
	request        *bufio.Buffer
	response       *bufio.Buffer
	tcpConnection  net.Conn
	conn           *Connection
	sync.Mutex
}

func (c *Connection) newContext() *Context {
	ctx := new(Context)
	ctx.conn = c
	ctx.request = bufio.NewBuffer(ctx.requestBuffer)
	ctx.response = bufio.NewBuffer(ctx.responseBuffer)

	return ctx
}

func (ctx *Context) Replay(message string) error {
	_, err := ctx.response.Write([]byte(message))
	return err
}

func (ctx *Context) Body() io.Reader {
	return ctx.request
}
