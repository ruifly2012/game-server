package main

import (
	"flag"
	"github.com/spf13/viper"
	"github.com/topfreegames/pitaya"
	"github.com/topfreegames/pitaya/logger"
	"github.com/topfreegames/pitaya/serialize/json"
	"os"
)

const (
	DEFAULT_ETCD_HOST = "localhost:2379"
	DEFAULT_NATS_HOST = "nats://localhost:4222"
)

func main() {

	bind := flag.String("bind", "0.0.0.0", "the server bind")
	port := flag.Int("port", 8080, "the server port")
	svType := flag.String("type", "web", "the server type")
	isFrontend := flag.Bool("frontend", false, "if server is frontend")

	flag.Parse()

	app, err := NewApp(*bind, *port)

	if err != nil {
		logger.Log.Fatal(err)
	}

	app.Init()

	defer pitaya.Shutdown()
	pitaya.SetSerializer(json.NewSerializer())

	ehost := os.Getenv("ETCD_HOST")
	if ehost == "" {
		ehost = DEFAULT_ETCD_HOST
	}

	nhost := os.Getenv("NATS_HOST")
	if nhost == "" {
		nhost = DEFAULT_NATS_HOST
	}

	config := viper.New()
	config.Set("pitaya.cluster.sd.etcd.endpoints", ehost)
	config.Set("pitaya.cluster.rpc.server.nats.connect", nhost)
	config.Set("pitaya.cluster.rpc.client.nats.connect", nhost)
	config.Set("pitaya.metrics.prometheus.enabled", true)
	config.Set("pitaya.handler.messages.compression", false)

	pitaya.Configure(*isFrontend, *svType, pitaya.Cluster, map[string]string{}, config)
	pitaya.Start()
}
