package main

import (
	"fmt"

	"github.com/labstack/echo"
	"github.com/vamsikartik01/Ethanhunt/services/sockets"
)

func hello(c echo.Context) error {
	c.String(200, "hello")
	return nil
}

func main() {
	fmt.Println("Hello World!")
	sockets.Initialize()
	ethan := echo.New()
	ethan.GET("/", hello)
	ethan.GET("/wsdevice", sockets.StartDeviceSocket)
	ethan.GET("/wsjames", sockets.StartJamesbondSocket)
	ethan.Logger.Fatal(ethan.Start(":2000"))
}
