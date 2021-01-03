package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {

	const myLatitude = 42.098733
	const myLongitude = -75.924118

	url := fmt.Sprintf("https://community-open-weather-map.p.rapidapi.com/forecast?&lat=%f&lon=%f", myLatitude, myLongitude)

	req, _ := http.NewRequest("GET", url, nil)

	req.Header.Add("x-rapidapi-key", "c4ed0a87d6mshc7315c833baa103p1c0e7ejsnf21eb690c64e")
	req.Header.Add("x-rapidapi-host", "community-open-weather-map.p.rapidapi.com")

	res, _ := http.DefaultClient.Do(req)

	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)

	fmt.Println(res)
	fmt.Println()
	fmt.Println(string(body))

}
