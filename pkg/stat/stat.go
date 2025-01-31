package stat

import (
	"github.com/inside-the-mirror/kratos/pkg/stat/prom"
)

// Stat interface.
type Stat interface {
	Timing(name string, time int64, extra ...string)
	Incr(name string, extra ...string) // name,ext...,code
	State(name string, val int64, extra ...string)
}

// default stat struct.
var (
	// http
	HTTPClient Stat = prom.HTTPClient
	HTTPServer Stat = prom.HTTPServer
	// storage
	Cache Stat = prom.LibClient
	DB    Stat = prom.LibClient
	// rpc
	RPCClient Stat = prom.RPCClient
	RPCServer Stat = prom.RPCServer
)
