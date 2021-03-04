package models

import (
	"crypto/md5"
	"errors"
	"fmt"
	"strings"
)

// ErrNoAvatarはAvatarインスタンスがアバターのURLを返すことができない
// 場合に発生するエラーです
var ErrNoAvatarURL = errors.New("chat: アバターのURLを取得できません。")

// Avatarはユーザーのプロフィール画像を表す型です。
type Avatar interface {
	// GetAvatarURLは指定されたクライアントのアバターのURLを返します。
	// 問題が発生した場合にはエラーを返します。特に、URLを取得できなかった
	// 場合にはErrNoAvatarURLを返します。
	AvatarURL(c *Client) (string, error)
	GravatarAvatarURL(c *Client) (string, error)
}

type AllAvatar struct{}

var UseAvatar AllAvatar

func (_ AllAvatar) AvatarURL(c *Client) (string, error) {
	url := c.User.AvatarURL
	if url != "" {
		return url, nil
	}
	return "", ErrNoAvatarURL
}

func (_ AllAvatar) GravatarAvatarURL(c *Client) (string, error) {
	email := c.User.Email
	if email != "" {
		md5 := md5.Sum([]byte(strings.ToLower(email)))
		return fmt.Sprintf("https://www.gravatar.com/avatar/%x", md5), nil
	}
	return "", ErrNoAvatarURL
}
