package mysql

import (
	"encoding/json"
	"log"
	"net/http"

	helper "github.com/vamsikartik01/Ethanhunt/helpers"
	"github.com/vamsikartik01/Ethanhunt/models"
)

func GetLocations(name string) ([]models.Location, error) {
	var Locations []models.Location

	url := "https://api.openweathermap.org/geo/1.0/direct?q=" + name + "&limit=5&appid=" + helper.Config.WeatherApiKey

	resp, err := http.Get(url)
	if err != nil {
		log.Println("Error making HTTP request:", err)
		return Locations, err
	}
	defer resp.Body.Close()

	if err := json.NewDecoder(resp.Body).Decode(&Locations); err != nil {
		log.Println("Error decoding JSON:", err)
		return Locations, err
	}

	return Locations, nil
}

func SetLocation(name string, lat string, lon string, country string, state string, accountSid int) error {
	query := "UPDATE Preferences SET name = ?, lat = ?, lon = ?, state = ?, city = ? WHERE accountSid = ?"

	result, err := db.Exec(query, name, lat, lon, country, state, accountSid)
	if err != nil {
		log.Println("error inserting into db, err : ", err)
		return err
	}

	affRows, err := result.RowsAffected()
	if err != nil {
		log.Println("error fetching affected rows, err : ", err)
		return err
	}
	log.Println("inserted %d row(s) into Hubs", affRows)
	return nil
}

func GetLocation(accountSid int) (models.LocationResp, error) {
	var location models.LocationResp
	rows, err := db.Query("select name, lat, lon, state, city from Preferences where accountSid = ?", accountSid)
	if err != nil {
		log.Println("Error Fetching rooms data")
		return location, err
	}
	defer rows.Close()

	for rows.Next() {
		err := rows.Scan(&location.Name, &location.Lat, &location.Lon, &location.Country, &location.State)
		if err != nil {
			log.Println("Error unmarshalling hubs.")
			return location, err
		}
	}

	if err := rows.Err(); err != nil {
		log.Println("Error processing rows: %v", err)
		return location, err
	}
	return location, nil
}
