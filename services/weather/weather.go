package weather

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	helper "github.com/vamsikartik01/Ethanhunt/helpers"
	"github.com/vamsikartik01/Ethanhunt/models"
	"github.com/vamsikartik01/Ethanhunt/services/mysql"
)

var jamesResponse models.JamesResponse
var lastTime int64 = 0

func GetWeatherData(accountSid int) (models.JamesResponse, error) {
	currTime := time.Now().Unix()
	log.Println("current Time: ", currTime, " last time: ", lastTime, " diff: ", currTime-lastTime)
	if currTime-lastTime < 3000 {
		log.Println("returnig cached weather")
		return jamesResponse, nil
	}

	location, err := mysql.GetLocation(accountSid)
	if err != nil {
		log.Printf("Error Fetching location Info")
		return jamesResponse, err
	}

	log.Println("locatino - ", location)

	url := "https://api.openweathermap.org/data/2.5/weather?lat=" + location.Lat + "&lon=" + location.Lon + "&units=metric&appid=" + helper.Config.WeatherApiKey
	log.Println("hitting weather api - ", url)
	response, err := http.Get(url)
	if err != nil {
		log.Printf("Error Fetching Weather Info")
		return jamesResponse, err
	}
	defer response.Body.Close()

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Printf("Error reading response body:", err)
		return jamesResponse, err
	}

	var weatherData models.WeatherResponse
	err = json.Unmarshal(body, &weatherData)
	if err != nil {
		log.Printf("Error Decoding Weather Info")
		return jamesResponse, err
	}

	jamesResponse.Name = weatherData.Name
	jamesResponse.Main = weatherData.Weather[0].Main
	jamesResponse.Description = weatherData.Weather[0].Description
	jamesResponse.Icon = weatherData.Weather[0].Icon
	jamesResponse.Temp = weatherData.Main.Temp
	jamesResponse.TempMin = weatherData.Main.TempMin
	jamesResponse.TempMax = weatherData.Main.TempMax
	jamesResponse.FeelsLike = weatherData.Main.FeelsLike
	jamesResponse.Speed = weatherData.Wind.Speed
	jamesResponse.Humidity = weatherData.Main.Humidity
	jamesResponse.Visibility = weatherData.Visibility
	jamesResponse.Pressure = weatherData.Main.Pressure
	jamesResponse.AirQuality = 0
	jamesResponse.State = location.State
	jamesResponse.Country = location.Country

	log.Printf("weather data %v", jamesResponse)
	lastTime = currTime
	return jamesResponse, nil
}
