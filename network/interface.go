package network

import (
	"io"
)

type IConnect interface {
	io.ReadWriteCloser
	GetLocalAddress() string
	GetRemoteAddress() string
}
