# 概要
GAE + GOで簡単なREST APIを作ってみた

# 構築手順
※ Go環境が整っている前提
```
$ mkdir $GOPATH/src/github.com/ahyataro/
$ cd $GOPATH/src/github.com/ahyataro/
$ git clone git@github.com:ahyataro/gae_go_rest_api.git
$ cd $GOPATH/src/github.com/ahyataro/gae_go_rest_api

# vendoring
$ dep init

# ローカル起動
$ dev_appserver.py app/app.yaml
```

# GAEへのデプロイコマンド
```
goapp deploy -application #{project_id] --version #{version} app
```