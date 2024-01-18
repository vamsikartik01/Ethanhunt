package mysql

import (
	"log"
)

func GetNote(accountSid int) (string, error) {
	var Note string
	rows, err := db.Query("select note from Preferences where accountSid = ?", accountSid)
	if err != nil {
		log.Println("Error Fetching Note data")
		return Note, err
	}
	defer rows.Close()

	if rows.Next() {
		err = rows.Scan(&Note)
		if err != nil {
			log.Println("Error unmarshalling Note.")
			return Note, err
		}
	}

	log.Println(Note)

	return Note, nil
}

func AddNote(note string, accountSid int) error {
	query := "UPDATE Preferences SET note = ? WHERE accountSid = ?"

	result, err := db.Exec(query, note, accountSid)
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
