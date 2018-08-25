package services

import (
	"github.com/topfreegames/pitaya"
	"github.com/topfreegames/pitaya/component"
)

type (
	// NNgame represents a component that contains a bundle of nngame related handler
	// like Join/Message
	NNgame struct {
		component.Base
		group *pitaya.Group
	}
)

// NewGame returns a new nngame
func NewGame() *NNgame {
	return &NNgame{
		group: pitaya.NewGroup("nngame"),
	}
}

// Init runs on service initialization
func (n *NNgame) Init() {}

// AfterInit component lifetime callback
func (n *NNgame) AfterInit() {
}
