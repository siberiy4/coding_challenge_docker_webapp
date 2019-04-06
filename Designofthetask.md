道筋だけ考えておく

Goを書いたことはないのでA Tour of Goを少しやって、

https://www.wantedly.com/companies/wantedly/post_articles/115474

DBって書いてるし、目指すはこれなんだろうなぁ(知らんけど)
## 1

https://fisproject.jp/2018/09/translates-grpc-into-rest-json-api-with-go/

こっちかな
https://qiita.com/sys_cat/items/80f1d8adf7eea31010a6


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
        Name: "Hello World!!"
    }
    return c.JSON(http.StatusOK, user)
}

で行けるんじゃないかなー
多分他の人は一日で終わるだろう


## 2
GIPでアクセスできる場所にコンテナおいて、port80でcurl
つまりファイアウォールの中の部サバは

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