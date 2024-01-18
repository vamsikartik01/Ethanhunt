package routes

import (
	"log"
	"net/http"

	"github.com/labstack/echo"
	"github.com/vamsikartik01/Ethanhunt/models"
	"github.com/vamsikartik01/Ethanhunt/services/mysql"
	"github.com/vamsikartik01/Ethanhunt/services/sockets"
	"github.com/vamsikartik01/Ethanhunt/services/weather"
)

func getAccountSid(c echo.Context) (int, error) {
	userClaims, ok := c.Get("user").(*models.JwtClaims)
	if !ok {
		return -1, echo.NewHTTPError(http.StatusInternalServerError, "Failed to retrieve user claims")
	}
	return userClaims.UserId, nil
}

func hello(c echo.Context) error {
	c.String(200, "hello")
	return nil
}

func GetWeather(c echo.Context) error {
	accountSid, err := getAccountSid(c)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Failed to retrieve user claims from context")
	}
	weatherInfo, err := weather.GetWeatherData(accountSid)
	if err != nil {
		log.Println("error fetching weather data-", err)
	}
	c.JSON(200, weatherInfo)
	return nil
}

func GetRooms(c echo.Context) error {
	accountSid, err := getAccountSid(c)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Failed to retrieve user claims from context")
	}
	rooms, err := mysql.GetRooms(accountSid)
	if err != nil {
		log.Println("error fetching rooms -", err)
	}
	c.JSON(200, rooms)
	return nil
}

func AddRoom(c echo.Context) error {
	accountSid, err := getAccountSid(c)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Failed to retrieve user claims from context")
	}
	name := c.QueryParam("name")
	err = mysql.AddRoom(name, accountSid)
	if err != nil {
		log.Println("error addinhg room -", err)
	}
	c.String(200, "true")
	return nil
}

func EditRoom(c echo.Context) error {
	accountSid, err := getAccountSid(c)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Failed to retrieve user claims from context")
	}
	name := c.QueryParam("name")
	id := c.QueryParam("id")
	err = mysql.UpdateRoom(name, id, accountSid)
	if err != nil {
		log.Println("error editing toom-", err)
	}
	c.String(200, "true")
	return nil
}

func DeleteRoom(c echo.Context) error {
	accountSid, err := getAccountSid(c)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Failed to retrieve user claims from context")
	}
	id := c.QueryParam("id")
	err = mysql.DeleteRoom(id, accountSid)
	if err != nil {
		log.Println("error deleteing room-", err)
	}
	c.String(200, "true")
	return nil
}

func GetHubs(c echo.Context) error {
	accountSid, err := getAccountSid(c)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Failed to retrieve user claims from context")
	}
	rooms, err := mysql.GetHubs(accountSid)
	if err != nil {
		log.Println("error fetching hubs - ", err)
	}
	c.JSON(200, rooms)
	return nil
}

func AddHub(c echo.Context) error {
	accountSid, err := getAccountSid(c)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Failed to retrieve user claims from context")
	}
	name := c.QueryParam("name")
	refId := c.QueryParam("refid")
	err = mysql.AddHub(name, accountSid, refId)
	if err != nil {
		log.Println("error adding hubs - ", err)
	}
	c.String(200, "true")
	return nil
}

func EditHub(c echo.Context) error {
	accountSid, err := getAccountSid(c)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Failed to retrieve user claims from context")
	}
	name := c.QueryParam("name")
	id := c.QueryParam("id")
	refId := c.QueryParam("refid")
	err = mysql.UpdateHub(name, id, refId, accountSid)
	if err != nil {
		log.Println("error editing hubs - ", err)
	}
	c.String(200, "true")
	return nil
}

func DeleteHub(c echo.Context) error {
	accountSid, err := getAccountSid(c)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Failed to retrieve user claims from context")
	}
	id := c.QueryParam("id")
	err = mysql.DeleteHub(id, accountSid)
	if err != nil {
		log.Println("error deleting hubs -", err)
	}
	c.String(200, "true")
	return nil
}

func GetDevices(c echo.Context) error {
	accountSid, err := getAccountSid(c)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Failed to retrieve user claims from context")
	}
	devices, err := mysql.GetDevices(accountSid)
	if err != nil {
		log.Println("error fetching devices -", err)
	}
	c.JSON(200, devices)
	return nil
}

func AddDevice(c echo.Context) error {
	accountSid, err := getAccountSid(c)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Failed to retrieve user claims from context")
	}
	name := c.QueryParam("name")
	roomId := c.QueryParam("roomid")
	hubId := c.QueryParam("hubid")
	mode := c.QueryParam("mode")
	Type := c.QueryParam("type")
	hubPort := c.QueryParam("hubport")
	err = mysql.AddDevice(name, accountSid, roomId, hubId, mode, Type, hubPort)
	if err != nil {
		log.Println("error adding device -", err)
	}
	c.String(200, "true")
	return nil
}

func EditDevice(c echo.Context) error {
	accountSid, err := getAccountSid(c)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Failed to retrieve user claims from context")
	}
	name := c.QueryParam("name")
	id := c.QueryParam("id")
	Type := c.QueryParam("type")
	hubPort := c.QueryParam("hubport")
	err = mysql.UpdateDevice(name, id, Type, hubPort, accountSid)
	if err != nil {
		log.Println("error editing device -", err)
	}
	c.String(200, "true")
	return nil
}

func SetFavorite(c echo.Context) error {
	accountSid, err := getAccountSid(c)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Failed to retrieve user claims from context")
	}
	id := c.QueryParam("id")
	value := c.QueryParam("value")
	err = mysql.SetFavorite(id, value, accountSid)
	if err != nil {
		log.Println("error setting favorite device-", err)
	}
	c.String(200, "true")
	return nil
}

func SetValue(c echo.Context) error {
	accountSid, err := getAccountSid(c)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Failed to retrieve user claims from context")
	}
	id := c.QueryParam("id")
	value := c.QueryParam("value")
	hubPort := c.QueryParam("hubport")
	hubRefId := c.QueryParam("hubrefid")
	err, stat := sockets.WriteDevice(hubRefId, hubPort, value)
	if err != nil {
		log.Println("error writing value to socket-", err)
	}
	var Value string
	if value == "1" {
		Value = "true"
	} else if value == "0" {
		Value = "false"
	}
	if !stat {
		err = mysql.SetValue(id, Value, accountSid)
		if err != nil {
			log.Println("error setting device value -", err)
		}
	}
	c.String(200, "true")
	return nil
}

func DeleteDevice(c echo.Context) error {
	accountSid, err := getAccountSid(c)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Failed to retrieve user claims from context")
	}
	id := c.QueryParam("id")
	err = mysql.DeleteDevice(id, accountSid)
	if err != nil {
		log.Println("error deleting device-", err)
	}
	c.String(200, "true")
	return nil
}

func GetNote(c echo.Context) error {
	accountSid, err := getAccountSid(c)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Failed to retrieve user claims from context")
	}
	note, err := mysql.GetNote(accountSid)
	if err != nil {
		log.Println("error fetching note-", err)
	}
	c.JSON(200, note)
	return nil
}

func SetNote(c echo.Context) error {
	accountSid, err := getAccountSid(c)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Failed to retrieve user claims from context")
	}
	note := c.QueryParam("note")
	err = mysql.AddNote(note, accountSid)
	if err != nil {
		log.Println("error setting note - ", err)
	}
	c.JSON(200, true)
	return nil
}

func GetLocations(c echo.Context) error {
	name := c.QueryParam("name")
	locations, err := mysql.GetLocations(name)
	if err != nil {
		log.Println("error fetching locations - ", err)
	}
	return c.JSON(200, locations)
}

func SetLocation(c echo.Context) error {
	accountSid, err := getAccountSid(c)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Failed to retrieve user claims from context")
	}
	name := c.QueryParam("name")
	lat := c.QueryParam("lat")
	lon := c.QueryParam("lon")
	country := c.QueryParam("country")
	state := c.QueryParam("state")
	err = mysql.SetLocation(name, lat, lon, country, state, accountSid)
	if err != nil {
		log.Println("error setting location - ", err)
	}

	return c.JSON(200, "true")
}
