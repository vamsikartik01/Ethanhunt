package sockets

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/websocket"
	"github.com/labstack/echo"
	helper "github.com/vamsikartik01/Ethanhunt/helpers"
	"github.com/vamsikartik01/Ethanhunt/models"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

var connectionsDevice = make(map[string]*websocket.Conn)
var connectionsJames = make(map[string]*websocket.Conn)

func Initialize() error {
	fmt.Println("Initialized sockets")
	return nil
}

func StartDeviceSocket(c echo.Context) error {
	conn, err := upgrader.Upgrade(c.Response(), c.Request(), nil)
	if err != nil {
		return err
	}
	defer conn.Close()

	connID := helper.GenerateSessionID()
	connectionsDevice[connID] = conn

	for {
		_, message, err := conn.ReadMessage()
		if err != nil {
			fmt.Println("error reading a message : ", err)
			break
		}

		var deviceMessage models.DeviceInfo
		if err := json.Unmarshal(message, &deviceMessage); err != nil {
			fmt.Println("Error unmarshalling the message : ", err)
			continue
		}

		fmt.Println("message: ", message)
		log.Println("message : ", deviceMessage.Name)

		conn.WriteMessage(websocket.TextMessage, message)

	}
	delete(connectionsDevice, connID)
	return nil
}

func StartJamesbondSocket(c echo.Context) error {
	conn, err := upgrader.Upgrade(c.Response(), c.Request(), nil)
	if err != nil {
		return err
	}
	defer conn.Close()

	connID := helper.GenerateSessionID()
	connectionsJames[connID] = conn

	for {
		_, message, err := conn.ReadMessage()
		if err != nil {
			fmt.Println("error reading a message : ", err)
			break
		}

		var deviceMessage models.DeviceInfo
		if err := json.Unmarshal(message, &deviceMessage); err != nil {
			fmt.Println("Error unmarshalling the message : ", err)
			continue
		}

		fmt.Println("message: ", message)
		log.Println("message : ", deviceMessage.Name)

		for id, con := range connectionsDevice {
			if id != connID {
				jamesWrite := models.DeviceInfo{
					Id:   "21",
					Name: "Name",
				}
				message, err = json.Marshal(jamesWrite)
				if err != nil {
					fmt.Println("Error marshalling the message : ", err)
					continue
				}
				con.WriteMessage(websocket.TextMessage, message)
			}
		}

		conn.WriteMessage(websocket.TextMessage, message)

	}
	delete(connectionsJames, connID)
	return nil
}

func StartJamesSocket(c echo.Context) error {
	fmt.Println("Sockets started")
	return nil
}
