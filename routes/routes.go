package routes

import (
	"github.com/labstack/echo"
	"github.com/vamsikartik01/Ethanhunt/services/sockets"
)

func DefineRoutes(ethan *echo.Echo) {
	ethan.GET("/", hello)
	ethan.GET("/wsdevice", sockets.StartDeviceSocket)

	ethan.GET("/room", GetRooms)
	ethan.POST("/room", AddRoom)
	ethan.PUT("/room", EditRoom)
	ethan.DELETE("/room", DeleteRoom)

	ethan.GET("/device", GetDevices)
	ethan.POST("/device", AddDevice)
	ethan.PUT("/device", EditDevice)
	ethan.PUT("/setfavorite", SetFavorite)
	ethan.PUT("/setvalue", SetValue)
	ethan.DELETE("/device", DeleteDevice)

	ethan.GET("/hub", GetHubs)
	ethan.POST("/hub", AddHub)
	ethan.PUT("/hub", EditHub)
	ethan.DELETE("/hub", DeleteHub)

	ethan.GET("/weather", GetWeather)

	ethan.GET("/location", GetLocations)
	ethan.PUT("/location", SetLocation)

	ethan.GET("/note", GetNote)
	ethan.PUT("/note", SetNote)
}
