package services

import (
	"context"
	"game-server/internal/protocol"
	"github.com/topfreegames/pitaya"
	"github.com/topfreegames/pitaya/component"
	"github.com/topfreegames/pitaya/logger"
)

type (
	// MJgame represents a component that contains a bundle of mjgame related handler
	// like Join/Message
	MJgame struct {
		component.Base
		group *pitaya.Group
	}
)

// NewGame returns a new mjgame
func NewGame() *MJgame {
	return &MJgame{
		group: pitaya.NewGroup("mjgame"),
	}
}

// Init runs on service initialization
func (m *MJgame) Init() {}

// AfterInit component lifetime callback
func (m *MJgame) AfterInit() {
}

// CreateRoom
func (m *MJgame) CreateRoom(ctx context.Context, req *protocol.CCreatePrivateRoom) (*protocol.JoinResponse, error) {

	logger.Log.Debugf("创建麻将房间")

	return &protocol.JoinResponse{Result: "ok"}, nil
}

// EnterRoom
func (m *MJgame) EntryRoom(ctx context.Context, req *protocol.CEnterSocialRoom) (*protocol.JoinResponse, error) {

	logger.Log.Debugf("加入麻将房间")

	return &protocol.JoinResponse{Result: "ok"}, nil
}

// Ready
func (m *MJgame) Ready(ctx context.Context) (*protocol.JoinResponse, error) {

	logger.Log.Debugf("玩家准备")

	return &protocol.JoinResponse{Result: "ok"}, nil
}

// Dissmiss
func (m *MJgame) Dissmiss(ctx context.Context) (*protocol.JoinResponse, error) {

	logger.Log.Debugf("玩家解散投票")

	return &protocol.JoinResponse{Result: "ok"}, nil
}

// Hu
func (m *MJgame) Hu(ctx context.Context) (*protocol.JoinResponse, error) {

	logger.Log.Debugf("玩家胡牌")

	return &protocol.JoinResponse{Result: "ok"}, nil
}

// Discard
func (m *MJgame) Discard(ctx context.Context) (*protocol.JoinResponse, error) {

	logger.Log.Debugf("玩家出牌")

	return &protocol.JoinResponse{Result: "ok"}, nil
}

// Operate
func (m *MJgame) Operate(ctx context.Context) (*protocol.JoinResponse, error) {

	logger.Log.Debugf("玩家操作")

	return &protocol.JoinResponse{Result: "ok"}, nil
}
