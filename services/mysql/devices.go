package mysql

import (
	"log"

	"github.com/vamsikartik01/Ethanhunt/models"
)

func GetDevices(accountSid int) ([]models.Devices, error) {
	var devices []models.Devices
	rows, err := db.Query("select * from Devices where accountSid = ?", accountSid)
	if err != nil {
		log.Println("Error Fetching Devices data")
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var device models.Devices
		err := rows.Scan(&device.Id, &device.Name, &device.Mode, &device.Status, &device.Type, &device.Value, &device.IsFavorite, &device.HubPort, &device.HubId, &device.RoomId, &device.AccountSid)
		if err != nil {
			log.Println("Error unmarshalling Devices.")
			return nil, err
		}
		devices = append(devices, device)
	}

	if err := rows.Err(); err != nil {
		log.Println("Error processing Devices: -", err)
		return nil, err
	}
	return devices, nil
}

func AddDevice(name string, accountSid int, roomId string, hubId string, mode string, Type string, hubPort string) error {
	query := "insert into Devices (name, mode, hubId, accountSid, roomId, type, hubPort) values (?, ?, ?, ?, ?, ?, ?)"

	result, err := db.Exec(query, name, mode, hubId, accountSid, roomId, Type, hubPort)
	if err != nil {
		log.Println("error inserting into db, err : -", err)
		return err
	}

	affRows, err := result.RowsAffected()
	if err != nil {
		log.Println("error fetching affected rows, err : -", err)
		return err
	}
	log.Println("inserted %d row(s) into Hubs", affRows)
	return nil
}

func UpdateDevice(name string, id string, Type string, hubPort string, accountSid int) error {
	query := "update Devices set name = ?, type = ?, hubPort = ? where id = ? and accountSid = ?"

	result, err := db.Exec(query, name, Type, hubPort, id, accountSid)
	if err != nil {
		log.Println("error updating row in Devices table, err : -", err)
		return err
	}

	affRows, err := result.RowsAffected()
	if err != nil {
		log.Println("error fetching affected rows, err : -", err)
		return err
	}

	log.Println("inserted %d row(s) into Devices", affRows)
	return nil
}

func DeleteDevice(id string, accountSid int) error {
	query := "delete from Devices where id = ? and accountSid = ?"

	result, err := db.Exec(query, id, accountSid)
	if err != nil {
		log.Println("error deleting rows in Devices table, err : -", err)
		return err
	}

	affRows, err := result.RowsAffected()
	if err != nil {
		log.Println("error fetching affected rows, err : -", err)
		return err
	}

	log.Println("inserted %d row(s) into Devices", affRows)
	return nil
}

func SetFavorite(id string, value string, accountSid int) error {
	var query string
	if value == "true" {
		query = "update Devices set isFavorite = true where id = ? and accountSid = ?"
	} else {
		query = "update Devices set isFavorite = false where id = ? and accountSid = ?"
	}

	result, err := db.Exec(query, id, accountSid)
	if err != nil {
		log.Println("error updating row in Devices table, err : -", err)
		return err
	}

	affRows, err := result.RowsAffected()
	if err != nil {
		log.Println("error fetching affected rows, err : -", err)
		return err
	}

	log.Println("inserted %d row(s) into Devices", affRows)
	return nil
}

func SetValue(id string, value string, accountSid int) error {
	var query string
	if value == "true" {
		query = "update Devices set value = true where id = ? and accountSid = ?"
	} else {
		query = "update Devices set value = false where id = ? and accountSid = ?"
	}

	result, err := db.Exec(query, id, accountSid)
	if err != nil {
		log.Println("error updating row in Devices table, err : -", err)
		return err
	}

	affRows, err := result.RowsAffected()
	if err != nil {
		log.Println("error fetching affected rows, err : -", err)
		return err
	}

	log.Println("inserted %d row(s) into Devices", affRows)
	return nil
}
