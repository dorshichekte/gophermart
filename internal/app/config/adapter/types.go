package config

import (
	"time"
)

type HTTPAdapter struct {
	Server HTTPServer
	Router Router
}

type HTTPServer struct {
	Address           string
	ReadHeaderTimeout time.Duration
	WriteTimeout      time.Duration
	ReadTimeout       time.Duration
	ShutdownTimeout   time.Duration
}

type Router struct {
	Shutdown time.Duration
	Timeout  time.Duration
}
