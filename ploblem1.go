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
