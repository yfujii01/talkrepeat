# talkrepeat

open jtalkを繰り返し呼び出す

## 準備

- インストール
```
$ brew install open-jtalk
$ go get github.com/yfujii01/talkrepeat
```

- 設定
```
$ export TALK_REPEAT_DIC=/usr/local/Cellar/open-jtalk/1.10_1/dic
$ export TALK_REPEAT_VOICE=/usr/local/Cellar/open-jtalk/1.10_1/voice/mei/mei_normal.htsvoice
$ export TALK_REPEAT_PLAY=afplay
```

- 設定(raspberry pi)
```
$ export TALK_REPEAT_DIC=/var/lib/mecab/dic/open-jtalk/naist-jdic
$ export TALK_REPEAT_VOICE=/usr/share/hts-voice/mei/mei_normal.htsvoice
$ export TALK_REPEAT_PLAY=aplay
```

- 更新
```
$ go get -u github.com/yfujii01/talkrepeat
```

- 削除
```
$ rm -rf $GOPATH/src/github.com/yfujii01/talkrepeat
$ rm $GOPATH/bin/talkrepeat
```

## 実行

```
$ talkrepeat
```