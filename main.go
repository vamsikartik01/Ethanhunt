package main

import (
	"fmt"
	"log"

	"github.com/labstack/echo"
	"github.com/vamsikartik01/Ethanhunt/services/mysql"
	"github.com/vamsikartik01/Ethanhunt/services/sockets"
)

func hello(c echo.Context) error {
	c.String(200, "hello")
	return nil
}

func get_rooms(c echo.Context) error {
	accountSid := c.Param("accountsid")
	log.Println("AccountSid : ", accountSid)
	rooms, err := mysql.GetRooms(accountSid)
	if err != nil {
		log.Println("error %v", err)
	}
	c.JSON(200, rooms)
	return nil
}

func add_room(c echo.Context) error {
	accountSid := c.Param("accountsid")
	name := c.QueryParam("name")
	log.Println("AccountSid : ", accountSid, " name: ", name)
	err := mysql.AddRoom(name, accountSid)
	if err != nil {
		log.Println("error %v", err)
	}
	c.String(200, "true")
	return nil
}

func edit_room(c echo.Context) error {
	accountSid := c.Param("accountsid")
	name := c.QueryParam("name")
	id := c.QueryParam("id")
	log.Println("AccountSid : ", accountSid, " name: ", name, " id: ", id)
	err := mysql.UpdateRoom(name, id)
	if err != nil {
		log.Println("error %v", err)
	}
	c.String(200, "true")
	return nil
}

func delete_room(c echo.Context) error {
	accountSid := c.Param("accountsid")
	id := c.QueryParam("id")
	log.Println("AccountSid : ", accountSid, " id: ", id)
	err := mysql.DeleteRoom(id)
	if err != nil {
		log.Println("error %v", err)
	}
	c.String(200, "true")
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
	ethan.GET("/getrooms/:accountsid", get_rooms)
	ethan.GET("/addroom/:accountsid", add_room)
	ethan.GET("/editroom/:accountsid", edit_room)
	ethan.GET("/deleteroom/:accountsid", delete_room)
	ethan.Logger.Fatal(ethan.Start(":2000"))
}
