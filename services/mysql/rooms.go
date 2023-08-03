package mysql

import (
	"log"

	"github.com/vamsikartik01/Ethanhunt/models"
)

func GetRooms(accountSid string) ([]models.Rooms, error) {
	var rooms []models.Rooms
	rows, err := db.Query("select * from Rooms where accountSid = ?", accountSid)
	if err != nil {
		log.Println("Error Fetching rooms data")
		return nil, err
	}
	defer rows.Close()

	if rows != nil {
		for rows.Next() {
			var room models.Rooms
			err := rows.Scan(&room.Id, &room.Name, &room.AccountSid)
			if err != nil {
				log.Println("Error unmarshalling rooms.")
				return nil, err
			}
			log.Println(room.Id, room.Name, room.AccountSid)
			rooms = append(rooms, room)
		}

		if err := rows.Err(); err != nil {
			log.Fatalf("Error processing rows: %v", err)
			return nil, err
		}
	}

	return rooms, nil
}

func AddRoom(name string, accountSid string) error {
	query := "insert into Rooms (name, accountSid) values (?, ?)"

	result, err := db.Exec(query, name, accountSid)
	if err != nil {
		log.Fatalf("error inserting into db, err : %v", err)
		return err
	}

	affRows, err := result.RowsAffected()
	if err != nil {
		log.Fatalf("error fetching affected rows, err : %v", err)
		return err
	}
	log.Println("inserted %d row(s) into Rooms", affRows)
	return nil
}

func UpdateRoom(name string, id string) error {
	query := "update Rooms set name = ? where id = ?"

	result, err := db.Exec(query, name, id)
	if err != nil {
		log.Fatalf("error updating row in rooms table, err : %v", err)
		return err
	}

	affRows, err := result.RowsAffected()
	if err != nil {
		log.Fatalf("error fetching affected rows, err : %v", err)
		return err
	}

	log.Println("inserted %d row(s) into Rooms", affRows)
	return nil
}

func DeleteRoom(id string) error {
	query := "delete from Rooms where id = ?"

	result, err := db.Exec(query, id)
	if err != nil {
		log.Fatalf("error deleting rows in rooms table, err : %v", err)
		return err
	}

	affRows, err := result.RowsAffected()
	if err != nil {
		log.Fatalf("error fetching affected rows, err : %v", err)
		return err
	}

	log.Println("inserted %d row(s) into Rooms", affRows)
	return nil
}
