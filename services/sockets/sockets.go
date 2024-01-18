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
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
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

	connID := helper.GenerateSessionID()
	connectionsDevice[connID] = conn
	log.Println("connection establixhed ", connID)
	temp := true
	var newConnID string

	for {
		if _, found := connectionsDevice[connID]; !found {
			connectionsDevice[connID] = conn
		}

		log.Println("waiting for message")
		_, message, err := conn.ReadMessage()
		if err != nil {
			fmt.Println("error reading a message : ", err)
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
				log.Printf("error: %v", err)
			}
			break
		}

		var deviceMessage models.DeviceInfo
		if err := json.Unmarshal(message, &deviceMessage); err != nil {
			fmt.Println("Error unmarshalling the message : ", err)
			continue
		}

		if temp {
			if deviceMessage.RefId != "" {
				newConnID = deviceMessage.RefId
				connectionsDevice[newConnID] = conn
				delete(connectionsDevice, connID)
				connID = newConnID
				temp = false
			} else {
				break
			}
		}

		if &deviceMessage != nil && &deviceMessage.Ping != nil && deviceMessage.Ping == "pinging" {
			message, err = json.Marshal(deviceMessage)
			if err != nil {
				fmt.Println("Error marshalling the message : ", err)
			}
			conn.WriteMessage(websocket.TextMessage, message)
		}

		fmt.Println("connections %v", connectionsDevice)
		fmt.Println("message: ", string(message))
		// log.Println("message : ", deviceMessage.Id)
		// log.Println("reference id : ", deviceMessage.RefId)
		// log.Println("completed message progress")
	}
	log.Println("Closing the connection!")
	delete(connectionsDevice, connID)
	delete(connectionsDevice, newConnID)
	conn.Close()
	return nil
}

func WriteDevice(hubRefID string, hubPort string, value string) (error, bool) {
	var conn *websocket.Conn
	log.Println("connections %v", connectionsDevice)
	if _, found := connectionsDevice[hubRefID]; found {
		conn = connectionsDevice[hubRefID]
	} else {
		return nil, true
	}
	var deviceMessage models.DeviceInfo
	deviceMessage.HubPort = hubPort
	deviceMessage.Value = value
	message, err := json.Marshal(deviceMessage)
	if err != nil {
		fmt.Println("Error marshalling the message to device: ", err)
		return err, false
	}
	err = conn.WriteMessage(websocket.TextMessage, message)
	if err != nil {
		return err, false
	}
	return nil, false
}
