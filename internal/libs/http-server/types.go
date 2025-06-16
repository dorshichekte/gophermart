package http_server

import (
	"go.uber.org/zap"
	"net/http"
	"time"
)

type Server struct {
	logger *zap.Logger
	server *http.Server
	config Config
}

type Config struct {
	Address           string
	ReadHeaderTimeout time.Duration
	WriteTimeout      time.Duration
	ReadTimeout       time.Duration
	ShutdownTimeout   time.Duration
}
