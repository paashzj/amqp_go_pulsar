package amsar

import (
	"github.com/sirupsen/logrus"
)

type Config struct {
	PulsarConfig PulsarConfig
}

type PulsarConfig struct {
	Host     string
	HttpPort int
	TcpPort  int
}

func Run(config *Config, impl Server) error {
	logrus.Info("Server Started")
	return nil
}
