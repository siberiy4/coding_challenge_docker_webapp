道筋だけ考えておく

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


## 3
気力をだせ