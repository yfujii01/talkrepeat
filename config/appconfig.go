package config

import "os"

// 構造体定義
type Config struct {
	Dictionary string
	Voice      string
}

// 環境変数を読み込み構造体へ割当
func Read() (*Config) {

	// 構造体を定義
	c := new(Config)

	c.Dictionary = os.Getenv("TALK_REPEAT_DIC")
	c.Voice = os.Getenv("TALK_REPEAT_VOICE")

	// 正常
	return c
}
