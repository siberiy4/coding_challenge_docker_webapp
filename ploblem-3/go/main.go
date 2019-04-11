package main

import (
    "github.com/jinzhu/gorm"
    "github.com/labstack/echo"
    "github.com/labstack/echo/middleware"
    _ "github.com/lib/pq"
    "log"
    "net/http"
)


//ここを合わせる
type Employee struct {
    Id     string `json:"id"`
    Name   string `json:"name"`
    Salary string `json:"salary"`
    Age    string `json:"age"`
}

type Users struct {
    Users []Employee `json:"employees"`
}

func main() {
    initMigrate()
    run()
}

func initMigrate() {
    db, err := gorm.Open("postgres", "host=localhost port=5432 user=wanted dbname=wanted-db password=wantedps sslmode=disable")
    if err != nil {
        log.Panic(err)
    }
    defer db.Close()
    db.AutoMigrate(&Employee{})
}

func run() {
    e := echo.New()
    e.Use(middleware.Logger())
    e.Use(middleware.Recover())

    e.GET("/employees", showAllUsers)
    e.GET("/employee/:id", showEmployee)
    e.PUT("/employee/:id", updateUsers)
    e.POST("/employee", newUsers)
    e.DELETE("/employee/:id", deleteEmployee)

    log.Fatal(e.Start(":8080"))
}

func showAllUsers(c echo.Context) error {
    db, err := gorm.Open("postgres", "host=localhost port=5432 user=wanted dbname=wanted-db password=wantedps sslmode=disable")
    defer db.Close()
    checkError(err)
    var employees []Employee
    db.Find(&employees)
    return c.JSON(http.StatusOK, employees)
}

func showEmployee(c echo.Context) error {
    db, err := gorm.Open("postgres", "host=localhost port=5432 user=wanted dbname=wanted-db password=wantedps sslmode=disable")
    checkError(err)
    defer db.Close()
    var e Employee
    id := c.Param("id")
    db.Where("id=?", id).Find(&e)
    return c.JSON(http.StatusOK, e)
}

func newUsers(c echo.Context) error {
    db, err := gorm.Open("postgres", "host=localhost port=5432 user=wanted dbname=wanted-db password=wantedps sslmode=disable")
    defer db.Close()
    checkError(err)

    employee := new(Employee)
    if err := c.Bind(employee); err != nil {
        return err
    }

    db.Create(&employee)
    return c.String(http.StatusOK, "OK")
}

func updateUsers(c echo.Context) error {
    var employee Employee
    db, err := gorm.Open("postgres", "host=localhost port=5432 user=wanted dbname=wanted-db password=wantedps sslmode=disable")
    checkError(err)
    defer db.Close()
    if err := c.Bind(&employee); err != nil {
        return err
    }
    paramId := c.Param("id")
    attrMap := map[string]interface{}{"name": employee.Name, "salary": employee.Salary, "age": employee.Age}
    db.Model(&employee).Where("id= ?", paramId).Updates(attrMap)
    return c.NoContent(http.StatusOK)
}

func deleteEmployee(c echo.Context) error {
    db, err := gorm.Open("postgres", "host=localhost port=5432 user=wanted dbname=wanted-db password=wantedps sslmode=disable")
    checkError(err)
    defer db.Close()
    var e Employee
    id := c.Param("id")
    db.Where("id=?", id).Find(&e).Delete(&e)

	// 削除した列のJSONを返すらしいので、ここを変更して、返すHTTPstatusも変更する？
    return c.JSON(http.StatusOK) 

}

func checkError(err error) {
    if err != nil {
        log.Panic(err)
    }
}