package main

import (
  "net/http"
  
  "github.com/labstack/echo/v4"
)

func main() {
  e := echo.New()

  e.GET("/", func (c echo.Context) error {
    return c.JSON(http.StatusOK, map[string]interface{}{
      "message": "Hello, world!",
    })
  })

  e.Start(":2100")
}
