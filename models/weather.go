package models

type WeatherResponse struct {
	Coord      Coord     `json:"coord"`
	Weather    []Weather `json:"weather"`
	Base       string    `json:"base"`
	Main       Main      `json:"main"`
	Visibility int       `json:"visibility"`
	Wind       Wind      `json:"wind"`
	Clouds     Clouds    `json:"clouds"`
	Dt         int64     `json:"dt"`
	Sys        Sys       `json:"sys"`
	Timezone   int       `json:"timezone"`
	ID         int       `json:"id"`
	Name       string    `json:"name"`
	Cod        int       `json:"cod"`
}

type Coord struct {
	Lon float64 `json:"lon"`
	Lat float64 `json:"lat"`
}

type Weather struct {
	ID          int    `json:"id"`
	Main        string `json:"main"`
	Description string `json:"description"`
	Icon        string `json:"icon"`
}

type Main struct {
	Temp      float64 `json:"temp"`
	FeelsLike float64 `json:"feels_like"`
	TempMin   float64 `json:"temp_min"`
	TempMax   float64 `json:"temp_max"`
	Pressure  int     `json:"pressure"`
	Humidity  int     `json:"humidity"`
	SeaLevel  int     `json:"sea_level"`
	GrndLevel int     `json:"grnd_level"`
}

type Wind struct {
	Speed float64 `json:"speed"`
	Deg   int     `json:"deg"`
	Gust  float64 `json:"gust"`
}

type Clouds struct {
	All int `json:"all"`
}

type Sys struct {
	Country string `json:"country"`
	Sunrise int64  `json:"sunrise"`
	Sunset  int64  `json:"sunset"`
}

type JamesResponse struct {
	Name        string  `json:"name"`
	Main        string  `json:"main"`
	Description string  `json:"description"`
	Icon        string  `json:"icon"`
	Temp        float64 `json:"temp"`
	TempMin     float64 `json:"temp_min"`
	TempMax     float64 `json:"temp_max"`
	FeelsLike   float64 `json:"feels_like"`
	Speed       float64 `json:"speed"`
	Humidity    int     `json:"humidity"`
	Visibility  int     `json:"visibility"`
	Pressure    int     `json:"pressure"`
	AirQuality  int     `json:"air_quality"`
	State       string  `json:"state"`
	Country     string  `json:"country"`
}
