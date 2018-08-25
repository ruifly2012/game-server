package main

import (
	"flag"
	"game-server/nngame-service/services"
	"os"
	"strings"

	"github.com/spf13/viper"
	"github.com/topfreegames/pitaya"
	"github.com/topfreegames/pitaya/component"
	"github.com/topfreegames/pitaya/serialize/json"
)

const (
	DEFAULT_ETCD_HOST = "localhost:2379"
	DEFAULT_NATS_HOST = "nats://localhost:4222"
)

func configureBackend() {
	nngame := services.NewGame()
	pitaya.Register(nngame,
		component.WithName("nngame"),
		component.WithNameFunc(strings.ToLower),
	)

	pitaya.RegisterRemote(nngame,
		component.WithName("nngameremote"),
		component.WithNameFunc(strings.ToLower),
	)
}

func main() {
	//port := flag.Int("port", 3251, "the port to listen")
	svType := flag.String("type", "nngame", "the server type")
	isFrontend := flag.Bool("frontend", false, "if server is frontend")

	flag.Parse()

	defer pitaya.Shutdown()

	pitaya.SetSerializer(json.NewSerializer())

	configureBackend()

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

	config.Set("pitaya.handler.messages.compression", false)

	pitaya.Configure(*isFrontend, *svType, pitaya.Cluster, map[string]string{}, config)
	pitaya.Start()
}
