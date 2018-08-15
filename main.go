package main

import (
	"os/exec"
	"os"
	"fmt"
	"bufio"
	"errors"
	"talkrepeat/config"
)

func main() {

	//configを読み込む
	c := config.Read()

	//辞書フォルダ存在確認
	if _, err := os.Stat(c.Dictionary); err != nil {
		fmt.Printf("config [dictionary] error (%v)\n", err)
		os.Exit(1)
	}

	//voiceファイル存在確認
	if _, err := os.Stat(c.Voice); err != nil {
		fmt.Printf("config [voice] error (%v)\n", err)
		os.Exit(1)
	}

	//exitが入力されるまで繰り返し処理を行う
	for {
		fmt.Print("(exitで終了)> ")

		//Scannerを使って一行読み
		scanner := bufio.NewScanner(os.Stdin)
		scanner.Scan()
		mes := scanner.Text()

		//入力文字がexitの場合は抜ける
		if mes == "exit" {
			break
		}

		//しゃべる
		say(mes, c)
	}

	//音声ファイルを削除する
	os.Remove("./out.wav")
}

func say(mes string, c *config.Config) {
	//テキストファイルを作成する
	FileWrite("./talkfile.txt", mes)

	//テキストファイルから音声ファイルを作成する
	exec.Command("open_jtalk",
		"-x", c.Dictionary,
		"-m", c.Voice,
		"-ow", "./out.wav", "./talkfile.txt").Output()

	//テキストファイルを削除する
	os.Remove("./talkfile.txt")

	//しゃべる
	exec.Command("afplay", "./out.wav").Start()
}

func FileWrite(filename string, mes string) (error) {

	if filename == ""{
		//log.Fatal("ファイル名が空白")
		return errors.New("ファイル名が空白です。")
	}

	file, err := os.OpenFile(filename, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		//log.Fatal(err)
		return errors.New("ファイル書き込みエラー")
	}
	defer file.Close()
	fmt.Fprintln(file, mes)

	return nil
}