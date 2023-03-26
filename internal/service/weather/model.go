package weather

type Wind struct {
	Speed float64 `json:"speed"`
	Deg   float64 `json:"deg"`
}

type Clouds struct {
	All int `json:"all"`
}

type Rain struct {
	Duration float64 `json:"1h"`
}

type Main struct {
	Temp      float64 `json:"temp"`
	FeelsLike float64 `json:"feels_like"`
	Pressure  int     `json:"pressure"`
	Humidity  int     `json:"humidity"`
	TempMin   float64 `json:"temp_min"`
	TempMax   float64 `json:"temp_max"`
}

type Weather struct {
	Id          int    `json:"id"`
	Main        string `json:"main"`
	Description string `json:"description"`
	Icon        string `json:"icon"`
}

type City struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

type Coord struct {
	Lon float64 `json:"lon"`
	Lat float64 `json:"lat"`
}

type CurrentWeatherResponse struct {
	Coord   `json:"coord"`
	Weather []Weather `json:"weather"`
	Main    `json:"main"`
	Wind    `json:"wind"`
	Rain    `json:"rain"`
	Clouds  `json:"clouds"`
	Dt      int    `json:"dt"`
	Id      int    `json:"id"`
	Name    string `json:"name"`
}

type ForecastResponse struct {
	City    `json:"city"`
	Coord   `json:"coord"`
	Country string `json:"country"`
	List    []struct {
		Dt      int `json:"dt"`
		Main    `json:"main"`
		Weather `json:"weather"`
		Clouds  `json:"clouds"`
		Wind    `json:"wind"`
	} `json:"list"`
}
