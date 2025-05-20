package options

import (
	"fmt"
	"net"
	"strconv"

	"github.com/ra1n6ow/fastgo/internal/apiserver"
	genericoptions "github.com/ra1n6ow/fastgo/pkg/options"
)

type ServerOptions struct {
	// MySQLOptions 字段在被 viper 解析时，会被解析到 mapstructure 对应的 mysql 字段下。
	MySQLOptions *genericoptions.MySQLOptions `json:"mysql" mapstructure:"mysql"`
	Addr         string                       `json:"addr" mapstructure:"addr"`
}

// NewServerOptions 创建带有默认值的 ServerOptions 实例.
func NewServerOptions() *ServerOptions {
	return &ServerOptions{
		MySQLOptions: genericoptions.NewMySQLOptions(),
		Addr:         "0.0.0.0:8087",
	}
}

// Validate 验证配置选项
func (o *ServerOptions) Validate() error {
	if err := o.MySQLOptions.Validate(); err != nil {
		return err
	}

	// 验证服务器地址
	if o.Addr == "" {
		return fmt.Errorf("addr is required")
	}

	// 检查地址格式是否为host:port
	_, portStr, err := net.SplitHostPort(o.Addr)
	if err != nil {
		return fmt.Errorf("invalid server address format '%s': %w", o.Addr, err)
	}

	// 验证端口是否为数字且在有效范围内
	port, err := strconv.Atoi(portStr)
	if err != nil || port < 1 || port > 65535 {
		return fmt.Errorf("invalid server port: %s", portStr)
	}

	return nil
}

func (o *ServerOptions) Config() (*apiserver.Config, error) {
	return &apiserver.Config{
		MySQLOptions: o.MySQLOptions,
		Addr:         o.Addr,
	}, nil
}
