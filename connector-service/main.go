package connector_service

import (
	"flag"
	"fmt"
	"game-server/connector-service/services"
	"os"
	"strings"

	"github.com/spf13/viper"
	"github.com/topfreegames/pitaya"
	"github.com/topfreegames/pitaya/acceptor"
	"github.com/topfreegames/pitaya/component"
	"github.com/topfreegames/pitaya/serialize/json"
)

const (
	DEFAULT_ETCD_HOST = "localhost:2379"
	DEFAULT_NATS_HOST = "nats://localhost:4222"
)

func configureFrontend(port int) {
	connector := services.NewConnector()
	pitaya.Register(connector,
		component.WithName("connector"),
		component.WithNameFunc(strings.ToLower),
	)
	pitaya.RegisterRemote(connector,
		component.WithName("connectorremote"),
		component.WithNameFunc(strings.ToLower),
	)

	ws := acceptor.NewWSAcceptor(fmt.Sprintf(":%d", port))
	pitaya.AddAcceptor(ws)
}

func main() {
	port := flag.Int("port", 3250, "the port to listen")
	svType := flag.String("type", "connector", "the server type")
	isFrontend := flag.Bool("frontend", true, "if server is frontend")

	flag.Parse()

	defer pitaya.Shutdown()

	pitaya.SetSerializer(json.NewSerializer())

	configureFrontend(*port)

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
