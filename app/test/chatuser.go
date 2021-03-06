package test

import gomniauthcommon "github.com/stretchr/gomniauth/common"

// 画像アップロードのstruct
type ChatUser interface {
	UniqueID() string
	AvatarURL() string
}

type chatUser struct {
	gomniauthcommon.User
	uniqueID string
}

func (u chatUser) UniqueID() string {
	return u.uniqueID
}
