package config

import (
	"fmt"
	"strings"
)

const (
	ModeProduction  = "production"
	ModeDevelopment = "development"
)

type Server struct {
	DomainName string `mapstructure:"domain_name"`
	Port       string `mapstructure:"port"`
	Mode       string `mapstructure:"server_mode"`
}

func (s *Server) SetMode(mode string) error {
	mode = strings.ToLower(mode)
	switch mode {
	case ModeProduction, ModeDevelopment:
		if mode != s.Mode {
			s.Mode = mode
		}
	default:
		return fmt.Errorf("invalid server mode: %s", mode)
	}
	return nil
}

// func (s *Server) GetMode() string {
// 	return s.ServerMode
// }

// func (s *Server) GetPort() string {
// 	return s.Port
// }
