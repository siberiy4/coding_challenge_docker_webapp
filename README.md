# coding_challenge_docker_webapp

## 概要
wantedlyの2019夏のインターンの応募課題

GoでAPI作って、dockerを使って公開しよう

## 課題1
ploblem1以下にDockerfileと main.goがあります。

echoを使いました。
Docker Hubにもあります。
https://hub.docker.com/r/yokanyukari/wantedsummer2019


## 課題2

AWSのEC2を借りてDockerをのせました。
http://52.192.48.108/

コンテナをポートフォワードして起動しました。
```
docker run --name wanted -d -p 80:8080   yokanyukari/wantedsummer2019
```

確認のコマンドは下のものになります。
```
curl -XGET -H 'Content-Type:application/json' http://52.192.48.108/
```

## 課題３

DBやdocker-composeを扱ったことがなかったのでやりかたを調べたかったが、5限までの授業が多く時間が足りずに終わってしまった。

docker-composeの仕組みが分かっていなかったのが一番の理由だと思う。
