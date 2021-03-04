package test

import (
	"crypto/md5"
	"errors"
	"fmt"
	"strings"
	"testing"
)

// ErrNoAvatarはAvatarインスタンスがアバターのURLを返すことができない
// 場合に発生するエラーです
var ErrNoAvatarURL = errors.New("chat: アバターのURLを取得できません。")

type Client struct {
	// socket *websocket.Conn //websocketの取得
	// send   chan *Message
	// room   *room
	User User
}

type User struct {
	// ID        int
	// UUID      string
	// Name      string
	Email string
	// PassWord  string
	AvatarURL string
	// CreatedAt time.Time
	// ApplyID   int
}

// Avatarはユーザーのプロフィール画像を表す型です。
type Avatar interface {
	// GetAvatarURLは指定されたクライアントのアバターのURLを返します。
	// 問題が発生した場合にはエラーを返します。特に、URLを取得できなかった
	// 場合にはErrNoAvatarURLを返します。
	AvatarURL(c *Client) (string, error)
}

type AuthAvatar struct{}

func (_ AuthAvatar) AvatarURL(c *Client) (string, error) {
	url := c.User.AvatarURL
	if url != "" {
		return url, nil
	}
	return "", ErrNoAvatarURL
}

func TestAuthAvatar(t *testing.T) {
	//値なしで確認
	var authAvatar AuthAvatar
	client := new(Client)
	url, err := authAvatar.AvatarURL(client)
	if err != ErrNoAvatarURL {
		t.Error("値が存在しない場合、AuthAvatar.GetAvatarURLは" +
			"ErrNoAvatarを返すべきです")
	}
	//値をセットします
	testURL := "http://url-to-avatar/"
	client.User.AvatarURL = testURL
	url, err = authAvatar.AvatarURL(client)
	if err != nil {
		t.Error("値が存在する場合、AuthAvatar.AvatarURLは" + "エラーをかえすべきではありません")
	} else {
		if url != testURL {
			t.Error("AuthAvatar.AvatarURLは正しいURLを返すべきです")
		}
	}
}

type GravatarAvatar struct{}

func (_ GravatarAvatar) AvatarURL(c *Client) (string, error) {
	email := c.User.Email
	if email != "" {
		md5 := md5.Sum([]byte(strings.ToLower(email)))
		return fmt.Sprintf("https://www.gravatar.com/avatar/%x", md5), nil
	}
	return "", ErrNoAvatarURL
}

const (
	email     = "keigo2356@gmail.com"
	avatarurl = "https://www.gravatar.com/avatar/1c55860804895a165a6bf5eca7c3cf3e"
)

func TestGravatarAvatar(t *testing.T) {
	var gravatarAvatar GravatarAvatar
	client := new(Client)
	client.User.Email = email
	url, err := gravatarAvatar.AvatarURL(client)
	if err != nil {
		t.Error("GravatarAvatar.AvatarURLはエラーを返すべきではありません")
	}
	if url != avatarurl {
		t.Errorf("GravatarAvatar.AvatarURLが%sという誤った値を返しました", url)
	}
}
