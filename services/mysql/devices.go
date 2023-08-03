package mysql

import (
	"log"

	"github.com/vamsikartik01/Ethanhunt/models"
)

func GetDevices(accountSid string) ([]models.Devices, error) {
	var devices []models.Devices
	rows, err := db.Query("select * from Devices where accountSid = ?", accountSid)
	if err != nil {
		log.Println("Error Fetching Devices data")
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var device models.Devices
		err := rows.Scan(&device.Id, &device.Name, &device.Mode, &device.Status, &device.HubId, &device.RoomId, &device.AccountSid, &device.Value)
		if err != nil {
			log.Println("Error unmarshalling Devices.")
			return nil, err
		}
		log.Println(device.Id, device.Name, device.AccountSid)
		devices = append(devices, device)
	}

	if err := rows.Err(); err != nil {
		log.Fatalf("Error processing Devices: %v", err)
		return nil, err
	}
	return devices, nil
}

func AddDevice(name string, accountSid string, roomId string, mode string, value string, hubId string) error {
	query := "insert into Devices (name, mode, value, hubId, accountSid, roomId) values (?, ?, ?, ?, ?, ?)"

	result, err := db.Exec(query, name, mode, value, hubId, accountSid, roomId)
	if err != nil {
		log.Fatalf("error inserting into db, err : %v", err)
		return err
	}

	affRows, err := result.RowsAffected()
	if err != nil {
		log.Fatalf("error fetching affected rows, err : %v", err)
		return err
	}
	log.Println("inserted %d row(s) into Hubs", affRows)
	return nil
}

func UpdateDevice(name string, id string) error {
	query := "update Devices set name = ? where id = ?"

	result, err := db.Exec(query, name, id)
	if err != nil {
		log.Fatalf("error updating row in Devices table, err : %v", err)
		return err
	}

	affRows, err := result.RowsAffected()
	if err != nil {
		log.Fatalf("error fetching affected rows, err : %v", err)
		return err
	}

	log.Println("inserted %d row(s) into Devices", affRows)
	return nil
}

func DeleteDevice(id string) error {
	query := "delete from Devices where id = ?"

	result, err := db.Exec(query, id)
	if err != nil {
		log.Fatalf("error deleting rows in Devices table, err : %v", err)
		return err
	}

	affRows, err := result.RowsAffected()
	if err != nil {
		log.Fatalf("error fetching affected rows, err : %v", err)
		return err
	}

	log.Println("inserted %d row(s) into Devices", affRows)
	return nil
}
