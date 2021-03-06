package models

import (
	"errors"
	"io/ioutil"
	"path/filepath"
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
	UploadAvatarURL(c *Client) (string, error)
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
	avatar_id := c.User.AvatarID
	if avatar_id != "" {
		return "https://www.gravatar.com/avatar/" + avatar_id, nil
	}
	return "", ErrNoAvatarURL
}

func (_ AllAvatar) UploadAvatarURL(c *Client) (string, error) {
	avatar_id := c.User.AvatarID
	if avatar_id != "" {
		if files, err := ioutil.ReadDir("app/views/avatars"); err == nil {
			for _, file := range files {
				if file.IsDir() {
					continue
				}
				if match, _ := filepath.Match(avatar_id+"*", file.Name()); match {
					return "/static/avatars/" + file.Name(), nil
				}
			}
		}
	}
	return "", ErrNoAvatarURL
}
