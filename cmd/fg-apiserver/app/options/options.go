package options

import (
	"github.com/ra1n6ow/fastgo/internal/apiserver"
	genericoptions "github.com/ra1n6ow/fastgo/pkg/options"
)

type ServerOptions struct {
	// MySQLOptions 字段在被 viper 解析时，会被解析到 mapstructure 对应的 mysql 字段下。
	MySQLOptions *genericoptions.MySQLOptions `json:"mysql" mapstructure:"mysql"`
}

// NewServerOptions 创建带有默认值的 ServerOptions 实例.
func NewServerOptions() *ServerOptions {
	return &ServerOptions{
		MySQLOptions: genericoptions.NewMySQLOptions(),
	}
}

// Validate 验证配置选项
func (s *ServerOptions) Validate() error {
	if err := s.MySQLOptions.Validate(); err != nil {
		return err
	}

	return nil
}

func (o *ServerOptions) Config() (*apiserver.Config, error) {
	return &apiserver.Config{
		MySQLOptions: o.MySQLOptions,
	}, nil
}
