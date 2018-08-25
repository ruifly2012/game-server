package services

import (
	"context"
	"game-server/connector-service/models"
	"game-server/internal/protocol"

	"github.com/topfreegames/pitaya/component"
)

// Connector struct
type Connector struct {
	component.Base
	players map[int64]*models.Player
}

// ConnectorRemote is a remote that will receive rpc's
type ConnectorRemote struct {
	component.Base
}

// NewConnector is a remote that will receive rpc's
func NewConnector() *Connector {
	return &Connector{
		players: map[int64]*models.Player{},
	}
}

// Init runs on service initialization
func (m *Connector) Init() {}

// AfterInit component lifetime callback
func (m *Connector) AfterInit() {}

// Entry is the entrypoint
func (c *Connector) Entry(ctx context.Context, req *protocol.LoginToGameServerRequest) (*protocol.JoinResponse, error) {

	return &protocol.JoinResponse{Result: "ok"}, nil
}
