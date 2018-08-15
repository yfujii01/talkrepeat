# talkrepeat

open jtalkを繰り返し呼び出す

## 事前準備

open jtalkをインストールする


go get github.com/yfujii01/talkrepeat


- 削除
rm -rf $GOPATH/src/github.com/yfujii01/talkrepeat
rm $GOPATH/bin/talkrepeat


- macの場合
```
brew install open-jtalk
```


## 設定ファイル修正

環境によって辞書ファイルとvoiceファイルのパスを修正する。
open_jtalkのインストール方法によって各々修正が必要。


export TALK_REPEAT_DIC=/usr/local/Cellar/open-jtalk/1.10_1/dic
export TALK_REPEAT_VOICE=/usr/local/Cellar/open-jtalk/1.10_1/voice/mei/mei_normal.htsvoice


- macの場合
```
{
  "dictionary": "/usr/local/Cellar/open-jtalk/1.10_1/dic",
  "voice": "/usr/local/Cellar/open-jtalk/1.10_1/voice/mei/mei_normal.htsvoice"
}
```

- raspiの場合
```
{
  "dictionary": "/var/lib/mecab/dic/open-jtalk/naist-jdic",
  "voice": "/usr/share/hts-voice/mei/mei_normal.htsvoice"
}
```

