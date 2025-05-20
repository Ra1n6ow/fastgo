package apiserver

import (
	"log/slog"

	genericoptions "github.com/ra1n6ow/fastgo/pkg/options"
)

// 用于存储应用相关的配置
type Config struct {
	MySQLOptions *genericoptions.MySQLOptions
}

type Server struct {
	cfg *Config
}

func (cfg *Config) NewServer() (*Server, error) {
	return &Server{cfg: cfg}, nil
}

func (s *Server) Run() error {
	slog.Info("Read MySQL host from config", "mysql.addr", s.cfg.MySQLOptions.Addr)

	// 阻塞防止进程退出
	select {}
}
