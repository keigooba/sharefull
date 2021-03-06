package models

import gomniauthcommon "github.com/stretchr/gomniauth/common"

// 画像アップロードのstruct
type ChatUser interface {
	UniqueID() string //これが呼ばれると構造体chatUserのuniqueIDを返す。
	AvatarURL() string
}

type chatUser struct {
	gomniauthcommon.User //これはインターフェース。この中にUserフィールド・メソッドAvatarURL()も含まれている。
	uniqueID string
}

func (u chatUser) UniqueID() string {
	return u.uniqueID
}
