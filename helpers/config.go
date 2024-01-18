package helper

import (
	"encoding/json"
	"log"
	"os"

	"github.com/vamsikartik01/Ethanhunt/models"
)

var Config *models.Config

func LoadConfig() error {
	filename := "config.json"
	file, err := os.Open(filename)
	if err != nil {
		log.Println("Error Loading Config file")
		return err
	}
	defer file.Close()

	err = json.NewDecoder(file).Decode(&Config)
	if err != nil {
		log.Println("Error Reading Json Object from file ", filename, " error -", err)
		return err
	}

	return nil
}
