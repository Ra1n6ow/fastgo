package apiserver

import (
	"log/slog"
	"net/http"

	"github.com/gin-gonic/gin"

	mw "github.com/ra1n6ow/fastgo/internal/pkg/middleware"
	genericoptions "github.com/ra1n6ow/fastgo/pkg/options"
)

// 用于存储应用相关的配置
type Config struct {
	MySQLOptions *genericoptions.MySQLOptions
	Addr         string
}

type Server struct {
	cfg *Config
	srv *http.Server
}

func (cfg *Config) NewServer() (*Server, error) {
	engine := gin.New()

	// Recovery 中间件，用来捕获任何 panic 恢复并返回 500 错误
	mws := []gin.HandlerFunc{gin.Recovery(), mw.NoCache, mw.Cors, mw.RequestID()}

	engine.Use(mws...)

	// 注册 404 Handler
	engine.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusNotFound, gin.H{"code": "PageNotFound", "message": "Page Not Found"})
	})

	// 注册健康检查 Handler
	engine.GET("/healthz", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"status": "OK"})
	})

	httpsrv := &http.Server{Addr: cfg.Addr, Handler: engine}

	return &Server{cfg: cfg, srv: httpsrv}, nil
}

func (s *Server) Run() error {
	slog.Info("Read MySQL host from config", "mysql.addr", s.cfg.MySQLOptions.Addr)

	if err := s.srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		return err
	}

	return nil
}
