# talkrepeat

open jtalkを繰り返し呼び出す

## 事前準備

open jtalkをインストールする

- macの場合
```
brew install open-jtalk
```


## 設定ファイル修正

環境によって辞書ファイルとvoiceファイルのパスを修正する。
open_jtalkのインストール方法によって各々修正が必要。

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

