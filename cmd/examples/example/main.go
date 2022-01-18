package main

import (
	"flag"
	"github.com/paashzj/amqp_go_pulsar/pkg/amsar"
	"os"
	"os/signal"
)

var pulsarHost = flag.String("pulsar_host", "localhost", "pulsar host")
var pulsarHttpPort = flag.Int("pulsar_http_port", 8080, "pulsar http port")
var pulsarTcpPort = flag.Int("pulsar_tcp_port", 6650, "pulsar tcp port")

func main() {
	flag.Parse()
	config := &amsar.Config{}
	config.PulsarConfig = amsar.PulsarConfig{}
	config.PulsarConfig.Host = *pulsarHost
	config.PulsarConfig.HttpPort = *pulsarHttpPort
	config.PulsarConfig.TcpPort = *pulsarTcpPort
	e := &ExampleAmsarImpl{}
	err := amsar.Run(config, e)
	if err != nil {
		panic(err)
	}
	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt)
	for {
		select {
		case <-interrupt:
			return
		}
	}
}
