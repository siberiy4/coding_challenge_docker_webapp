道筋と参考(コピぺしたサイト)


Goを書いたことはないのでA Tour of Goを少しやって、

https://www.wantedly.com/companies/wantedly/post_articles/115474

DBって書いてるし、目指すはこれなんだろうなぁ(知らんけど)

https://akirachiku.com/post/2018-08-11-go-net-http-api-server-6/
これとか参考にしたい人生だった。

## 1

https://fisproject.jp/2018/09/translates-grpc-into-rest-json-api-with-go/

こっちかな
https://qiita.com/sys_cat/items/80f1d8adf7eea31010a6


最後のjsonのところをコピペ改変した。

```
package main

import (
    "net/http"
    "github.com/labstack/echo"
)

type User struct {
    Message  string `json:"message"`
}

func main() {
    e := echo.New()
    e.GET("/", hello)
    e.Logger.Fatal(e.Start(":8080"))
}

func hello(c echo.Context) error {
    user := &User{
        Message: "Hello World!!",
    }
    return c.JSON(http.StatusOK, user)
}

```

で行けるんじゃないかなー
多分他の人は一日で終わるだろう


ploblem1.goに書いた

はまりどころは
20行目の `Message: "Hello World!!",` のさいごの `,` が必須であること。
感想：フレームワーク強い。
あとはalpineあたりにマウントしてハブに出も上げよう


### day2

ビルドしたバイナリを別のイメージにコピーするのなんだっけってずっとなっていたが、odan先輩が教えてくれた。
multi stage build
docker ちょっと調べたときにも見たし、サイバーでも聞いたはずなのになんで忘れるか・・・。
バイナリを乗せるimageは選定するらしいが、今回はalpineでいいと思う


参考：
https://qiita.com/ight/items/a9f69d4617f17e2c2f21
https://qiita.com/minamijoyo/items/711704e85b45ff5d6405

alpine はちょっと特殊なので
https://qiita.com/katoken-0215/items/f3a502fe0c2044709012


https://hub.docker.com/r/yokanyukari/wantedsummer2019

docker hubには登録した

## 2
GIPでアクセスできる場所にコンテナおいて、port80でcurl
つまりファイアウォールの中の部サバはダメ

PaaSを使ったことないので、ふつうのEC2借りてdockerの設定のところでなんやかんや出来たような

```
{
  "message": "Hello World!!"
}
```
を返そう。
コンテナは問1で作ったので、docker pullしたりして、
コンテナを起動、port フォワードすればできそう

### day2

https://qiita.com/tifa2chan/items/a58e34019d4f10097a4d

```
docker run --name wanted -d -p 80:8080   yokanyukari/wantedsummer2019
```
で終わり


## 3
気力をだせ


https://qiita.com/t0w4/items/e886a514559cdb295600



### day3

https://qiita.com/muroya2355/items/d48c384a4a82c7ed34ae

https://qiita.com/hiro9/items/e6e41ec822a7077c3568

https://qiita.com/mztnnrt/items/8f671b39269b68e9ef5d


これがGoでSQL触るやつ
https://qiita.com/mztnnrt/items/eaa4843b2bb3f7bdaecd


1~5限はしんどい