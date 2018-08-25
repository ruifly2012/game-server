package models

import "github.com/topfreegames/pitaya/session"

// Player player
type Player struct {
	uid  int64  // 用户ID
	head string // 头像地址
	name string // 玩家名字
	ip   string // ip地址
	sex  int    // 性别
	coin int64  // 房卡数量
}

func NewPlayer(s *session.Session, uid int64, name, head, ip string, sex int) *Player {
	p := &Player{
		uid:  uid,
		name: name,
		head: head,
		ip:   ip,
		sex:  sex,
		coin: 10,
	}

	return p
}
