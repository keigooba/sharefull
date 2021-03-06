package test

import (
	"errors"
	"testing"
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
}

type AllAvatar struct{}

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
		return "/static/avatars/" + avatar_id + ".jpg", nil
	}
	return "", ErrNoAvatarURL
}

func TestAuthAvatar(t *testing.T) {
	//値なしで確認
	var authAvatar AllAvatar
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

const (
	avatar_id = "1c55860804895a165a6bf5eca7c3cf3e"
	avatarurl = "https://www.gravatar.com/avatar/1c55860804895a165a6bf5eca7c3cf3e"
	u_avatarurl = "/static/avatars/1c55860804895a165a6bf5eca7c3cf3e.jpg"
)

func TestGravatarAvatar(t *testing.T) {
	var gravatarAvatar AllAvatar
	client := new(Client)
	client.User.AvatarID = avatar_id
	url, err := gravatarAvatar.GravatarAvatarURL(client)
	if err != nil {
		t.Error("GravatarAvatar.AvatarURLはエラーを返すべきではありません")
	}
	if url != avatarurl {
		t.Errorf("GravatarAvatar.AvatarURLが%sという誤った値を返しました", url)
	}
}

func TestUploadAvatar(t *testing.T) {
	var uploadAvatar AllAvatar
	client := new(Client)
	client.User.AvatarID = avatar_id
	url, err := uploadAvatar.UploadAvatarURL(client)
	if err != nil {
		t.Error("UploadAvatar.AvatarURLはエラーを返すべきではありません")
	}
	if url != u_avatarurl {
		t.Errorf("UploadAvatar.AvatarURLが%sという誤った値を返しました", url)
	}
}
