package main

import (
	"fmt"

	"github.com/labstack/echo"
	"github.com/vamsikartik01/Ethanhunt/services/mysql"
	"github.com/vamsikartik01/Ethanhunt/services/sockets"
)

func hello(c echo.Context) error {
	c.String(200, "hello")
	return nil
}

func get_rooms(c echo.Context) error {
	mysql.GetRooms()
	return nil
}

func main() {
	fmt.Println("Hello World!")
	sockets.Initialize()
	mysql.InitConnection()
	ethan := echo.New()
	ethan.GET("/", hello)
	ethan.GET("/wsdevice", sockets.StartDeviceSocket)
	ethan.GET("/wsjames", sockets.StartJamesbondSocket)
	ethan.GET("/getrooms", get_rooms)
	ethan.Logger.Fatal(ethan.Start(":2000"))
}
