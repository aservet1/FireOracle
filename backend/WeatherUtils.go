package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"regexp"
	"strings"

	"math"
	"math/rand"
	"time"
)

type WeatherInfo struct {
	temperature float64 // ie 37.98
	pressure    float64 // ie 1019
	humidity    float64 // ie 79
	wind_speed  float64 // ie 4.47, assuming it's in knots
	wind_degree float64 // ie 148 -- out of 360
	clouds      float64 // ie 15 -- 1-100
}

const temperature = 0
const pressure = 1
const humidity = 2
const wind_speed = 3
const wind_degree = 4
const clouds = 5

func getWeatherInfoStruct(wiSlice []float64) WeatherInfo {
	var wi WeatherInfo
	wi.temperature = wiSlice[temperature]
	wi.pressure = wiSlice[pressure]
	wi.clouds = wiSlice[clouds]
	wi.wind_degree = wiSlice[wind_degree]
	wi.wind_speed = wiSlice[wind_speed]
	wi.humidity = wiSlice[humidity]
	return wi
}

func Forecast(myCounty string) [][]float64 {

	myKey := "c4ed0a87d6mshc7315c833baa103p1c0e7ejsnf21eb690c64e"
	// url := fmt.Sprintf("https://community-open-weather-map.p.rapidapi.com/forecast?q=%s%%2Cus&units=imperial", myCounty)
	Url, _ := url.Parse("https://community-open-weather-map.p.rapidapi.com")
	Url.Path += "/forecast"
	parameters := url.Values{}
	parameters.Add("q", myCounty+",us")
	parameters.Add("units", myCounty+"impertial")
	Url.RawQuery = parameters.Encode()
	urL := Url.String()

	req, _ := http.NewRequest("GET", urL, nil)
	req.Header.Add("x-rapidapi-key", myKey)
	req.Header.Add("x-rapidapi-host", "community-open-weather-map.p.rapidapi.com")
	res, _ := http.DefaultClient.Do(req)
	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)

	umsh := make(map[string]interface{})
	json.Unmarshal([]byte(body), &umsh)
	list := umsh["list"].([]interface{})

	var wis [][]float64
	for i := 0; i < len(list); i++ {

		item := list[i].(map[string]interface{})

		main_info := item["main"].(map[string]interface{})
		wind := item["wind"].(map[string]interface{})
		cloudss := item["clouds"].(map[string]interface{})

		wi := make([]float64, 6)
		//wi.dt, _ = time.Parse(time.RFC3339, strings.Join(strings.Split(item["dt_txt"].(string)," "),"T")+".00Z")
		wi[temperature] = main_info["temp"].(float64)
		wi[pressure] = main_info["pressure"].(float64)
		wi[humidity] = main_info["humidity"].(float64)
		wi[wind_degree] = wind["deg"].(float64)
		wi[wind_speed] = wind["speed"].(float64)
		wi[clouds] = cloudss["all"].(float64)

		wis = append(wis, wi)
	}
	return wis
}

//private helper
func extract(label, s string) string {

	pattern := fmt.Sprintf("\"%s\":{[^}]*}", label)
	re := regexp.MustCompile(pattern)
	reres := re.FindAllStringSubmatch(s, -1)[0][0]
	return strings.ReplaceAll(reres, fmt.Sprintf("\"%s\":", label), "")

}

func Current(county string) []float64 {

	url := fmt.Sprintf("https://community-open-weather-map.p.rapidapi.com/weather?q=%s%%2Cus&lat=0&lon=0&callback=test&lang=null&units=imperial", county)

	req, _ := http.NewRequest("GET", url, nil)

	req.Header.Add("x-rapidapi-key", "c4ed0a87d6mshc7315c833baa103p1c0e7ejsnf21eb690c64e")
	req.Header.Add("x-rapidapi-host", "community-open-weather-map.p.rapidapi.com")

	res, _ := http.DefaultClient.Do(req)

	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)

	var main_info map[string]interface{}
	var wind_info map[string]interface{}
	var cloudss map[string]interface{}
	json.Unmarshal([]byte(extract("main", string(body))), &main_info)
	json.Unmarshal([]byte(extract("wind", string(body))), &wind_info)
	json.Unmarshal([]byte(extract("clouds", string(body))), &cloudss)

	wi := make([]float64, 6)
	//wi.dt = time.Now()
	wi[temperature] = main_info["temp"].(float64)
	wi[pressure] = main_info["pressure"].(float64)
	wi[humidity] = main_info["humidity"].(float64)
	wi[wind_degree] = wind_info["deg"].(float64)
	wi[wind_speed] = wind_info["speed"].(float64)
	wi[clouds] = cloudss["all"].(float64)

	return wi
}

func RandomWeatherInfo() []float64 { //assumed randomWeather has been called, so rand has been seeded
	// rand.Seed(time.Now().UnixNano())
	wi := make([]float64, 6)
	wi[temperature] = rand.Float64() * 99
	wi[pressure] = rand.Float64() + 990
	wi[humidity] = rand.Float64()
	wi[wind_speed] = rand.Float64() * 10
	wi[wind_degree] = rand.Float64() * 360
	wi[clouds] = rand.Float64() * 100
	return wi

}

func RandomHistoricalWeather(start, end time.Time) [][]float64 {
	rand.Seed(time.Now().UnixNano())
	
	if start.IsZero() {
		return nil
	}
	var days int
	if end.IsZero() {
		days = 1
	} else {
		days = int(math.Ceil(end.Sub(start).Hours()/24)) - 1
	}

	var weatherPerDay [][]float64
	for i := 0; i < days; i++ {
		weatherPerDay = append(weatherPerDay, RandomWeatherInfo())
	}
	return weatherPerDay
}
