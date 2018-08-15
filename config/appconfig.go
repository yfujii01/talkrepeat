package config

import (
	"io/ioutil"
	"encoding/json"
)

// 構造体定義
type Config struct {
	Dictionary string `json:"dictionary"`
	Voice      string `json:"voice"`
}

// configファイルを読み込み構造体へ割当
func Read(filename string) (*Config, error) {

	// 構造体を定義
	c := new(Config)

	// 設定ファイルを読み込む
	jsonString, err := ioutil.ReadFile(filename)
	if err != nil {
		// エラー
		return c, err
	}

	// 設定
	err = json.Unmarshal(jsonString, c)
	if err != nil {
		// エラー
		return c, err
	}

	// 正常
	return c, nil
}
