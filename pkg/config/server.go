package config

import (
	"strconv"
)

type Server struct {
	Addr string
	Port int
}

// ToAddr 拼接 Addr:Port
func (t *Server) ToAddr() string {
	return t.Addr + ":" + strconv.Itoa(t.Port)
}
