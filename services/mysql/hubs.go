package mysql

import (
	"log"

	"github.com/vamsikartik01/Ethanhunt/models"
)

func GetHubs(accountSid int) ([]models.Hubs, error) {
	var hubs []models.Hubs
	rows, err := db.Query("select * from Hubs where accountSid = ?", accountSid)
	if err != nil {
		log.Println("Error Fetching rooms data")
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var hub models.Hubs
		err := rows.Scan(&hub.Id, &hub.Name, &hub.AccountSid, &hub.RefId)
		if err != nil {
			log.Println("Error unmarshalling hubs.")
			return nil, err
		}
		hubs = append(hubs, hub)
	}

	if err := rows.Err(); err != nil {
		log.Println("Error processing rows: %v", err)
		return nil, err
	}
	return hubs, nil
}

func AddHub(name string, accountSid int, refId string) error {
	query := "insert into Hubs (name, accountSid, refId) values (?, ?, ?)"

	result, err := db.Exec(query, name, accountSid, refId)
	if err != nil {
		log.Println("error inserting into db, err : %v", err)
		return err
	}

	affRows, err := result.RowsAffected()
	if err != nil {
		log.Println("error fetching affected rows, err : %v", err)
		return err
	}
	log.Println("inserted %d row(s) into Hubs", affRows)
	return nil
}

func UpdateHub(name string, id string, refId string, accountSid int) error {
	query := "update Hubs set name = ?,refId = ? where id = ? and accountSid = ?"

	result, err := db.Exec(query, name, refId, id, accountSid)
	if err != nil {
		log.Println("error updating row in Hubs table, err : %v", err)
		return err
	}

	affRows, err := result.RowsAffected()
	if err != nil {
		log.Println("error fetching affected rows, err : %v", err)
		return err
	}

	log.Println("updated %d row(s) into Hubs", affRows)
	return nil
}

func DeleteHub(id string, accountSid int) error {
	query := "delete from Hubs where id = ? and accountSid = ?"

	result, err := db.Exec(query, id, accountSid)
	if err != nil {
		log.Println("error deleting rows in hubs table, err : %v", err)
		return err
	}

	affRows, err := result.RowsAffected()
	if err != nil {
		log.Println("error fetching affected rows, err : %v", err)
		return err
	}

	log.Println("inserted %d row(s) into Hubs", affRows)
	return nil
}
