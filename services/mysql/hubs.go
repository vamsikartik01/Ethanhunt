package mysql

import (
	"log"

	"github.com/vamsikartik01/Ethanhunt/models"
)

func GetHubs(accountSid string) ([]models.Hubs, error) {
	var hubs []models.Hubs
	rows, err := db.Query("select * from Hubs where accountSid = ?", accountSid)
	if err != nil {
		log.Println("Error Fetching rooms data")
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var hub models.Hubs
		err := rows.Scan(&hub.Id, &hub.Name, &hub.AccountSid, &hub.RoomId)
		if err != nil {
			log.Println("Error unmarshalling rooms.")
			return nil, err
		}
		log.Println(hub.Id, hub.Name, hub.AccountSid)
		hubs = append(hubs, hub)
	}

	if err := rows.Err(); err != nil {
		log.Fatalf("Error processing rows: %v", err)
		return nil, err
	}
	return hubs, nil
}

func AddHub(name string, accountSid string, roomId string) error {
	query := "insert into Hubs (name, accountSid, roomId) values (?, ?, ?)"

	result, err := db.Exec(query, name, accountSid, roomId)
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

func UpdateHub(name string, id string) error {
	query := "update Hubs set name = ? where id = ?"

	result, err := db.Exec(query, name, id)
	if err != nil {
		log.Fatalf("error updating row in Hubs table, err : %v", err)
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

func DeleteHub(id string) error {
	query := "delete from Hubs where id = ?"

	result, err := db.Exec(query, id)
	if err != nil {
		log.Fatalf("error deleting rows in hubs table, err : %v", err)
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
