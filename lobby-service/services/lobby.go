package services

import (
	"context"
	"fmt"
	"game-server/internal/protocol"

	"github.com/google/uuid"
	"github.com/topfreegames/pitaya"
	"github.com/topfreegames/pitaya/component"
	"github.com/topfreegames/pitaya/timer"
)

type (
	// Lobby represents a component that contains a bundle of lobby related handler
	// like Join/Message
	Lobby struct {
		component.Base
		group *pitaya.Group
		timer *timer.Timer
	}

	// UserMessage represents a message that user sent
	UserMessage struct {
		Name    string `json:"name"`
		Content string `json:"content"`
	}

	// NewUser message will be received when new user join lobby
	NewUser struct {
		Content string `json:"content"`
	}

	// AllMembers contains all members uid
	AllMembers struct {
		Members []string `json:"members"`
	}
)

// NewLobby returns a new lobby
func NewLobby() *Lobby {
	return &Lobby{
		group: pitaya.NewGroup("lobby"),
	}
}

// Init runs on service initialization
func (r *Lobby) Init() {}

// AfterInit component lifetime callback
func (r *Lobby) AfterInit() {
}

// Entry is the entrypoint
func (r *Lobby) Entry(ctx context.Context, msg []byte) (*protocol.JoinResponse, error) {
	fakeUID := uuid.New().String() // just use s.ID as uid !!!
	s := pitaya.GetSessionFromCtx(ctx)
	err := s.Bind(ctx, fakeUID) // binding session uid
	if err != nil {
		return nil, pitaya.Error(err, "RH-000", map[string]string{"failed": "bind"})
	}
	return &protocol.JoinResponse{Result: "ok"}, nil
}

// Join lobby
func (r *Lobby) Join(ctx context.Context) (*protocol.JoinResponse, error) {
	s := pitaya.GetSessionFromCtx(ctx)
	err := r.group.Add(s)
	if err != nil {
		return nil, err
	}
	s.Push("onMembers", &AllMembers{Members: r.group.Members()})
	r.group.Broadcast("onNewUser", &NewUser{Content: fmt.Sprintf("New user: %d", s.ID())})
	if err != nil {
		return nil, err
	}
	return &protocol.JoinResponse{Result: "success"}, nil
}

// Message sync last message to all members
func (r *Lobby) Message(ctx context.Context, msg *UserMessage) {
	err := r.group.Broadcast("onMessage", msg)
	if err != nil {
		fmt.Println("error broadcasting message", err)
	}
}
