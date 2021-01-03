package main

import (
	"database/sql"
	"errors"
	"fmt"
	"runtime"
	"sort"
	"strconv"
	"sync"
	"time"

	"github.com/cdipaolo/goml/base"
	"github.com/cdipaolo/goml/linear"
	_ "github.com/mattn/go-sqlite3"
)

// GetFires returns all the fires in the county specified
func GetFires(db *sql.DB, county string, year int) ([]time.Time, error) {
	sql := fmt.Sprintf("SELECT FIRE_YEAR, DISCOVERY_DOY FROM fires_%d WHERE COUNTY like ?;", year)
	stmt, err := db.Prepare(sql)
	if err != nil {
		fmt.Println(err)
		return nil, errors.New("Error while preparing query")
	}
	dates := make([]time.Time, 0)
	rows, err := stmt.Query(county)
	if err != nil {
		fmt.Println(err)
		return nil, errors.New("Error while executing query")
	}
	defer rows.Close()
	for rows.Next() {
		var (
			fireYearStr string
			fireDOYStr  string
		)
		if err := rows.Scan(&fireYearStr, &fireDOYStr); err != nil {
			fmt.Println(err)
			return dates, errors.New("Error while reading row")
		}
		if fireYearStr != "" && fireDOYStr != "" {
			fireYearInt, _ := strconv.Atoi(fireYearStr)
			fireDOYInt, _ := strconv.Atoi(fireDOYStr)
			fireTime := time.Date(fireYearInt, time.January, 1, 0, 0, 0, 0, time.UTC)
			fireTime = fireTime.AddDate(0, 0, fireDOYInt-1)
			dates = append(dates, fireTime)
		}
	}
	return dates, nil
}

func getObserved(county string, start, end int, parallel bool) ([]float64, time.Time, time.Time) {
	sqliteDB, err := sql.Open("sqlite3", "./project.sqlite")
	var first, last time.Time
	if err != nil {
		fmt.Println(err)
		return nil, first, last
	}
	defer sqliteDB.Close()
	datesSlice := NewParallelResultSlice()
	dateResultChan := make(chan []time.Time)
	go datesSlice.Listen(dateResultChan)
	if parallel {
		var wg sync.WaitGroup
		for year := start; year <= end; year++ {

			wg.Add(1)
			go func(year int) {
				defer wg.Done()
				dates, err := GetFires(sqliteDB, county, year)
				if err != nil {
					fmt.Println(err)
					return
				}
				datesSlice.AddDates(dates)
			}(year)
		}
		wg.Wait()
	} else {
		for year := start; year <= end; year++ {
			fmt.Println(year)
			dates, err := GetFires(sqliteDB, county, year)
			if err != nil {
				fmt.Println(err)
				continue
			}
			datesSlice.AddDates(dates)
		}
	}
	datesSlice.Kill()
	dates := <-dateResultChan
	zeroOne := make([]float64, 0)
	sort.Slice(dates, func(i, j int) bool {
		return dates[i].Before(dates[j])
	})
	for i := 0; i < len(dates)-1; i++ {
		zeroOne = append(zeroOne, 1)
		diff := (int(dates[i+1].Sub(dates[i]).Hours()) / 24) - 1
		for j := 0; j < diff; j++ {
			zeroOne = append(zeroOne, 0)
		}
		if i == len(dates)-2 {
			zeroOne = append(zeroOne, 1)
			break
		}
	}
	n := len(dates)
	if n >= 1 {
		first = dates[0]
		if n >= 2 {
			last = dates[n-1]
		}
	}
	return zeroOne, first, last
}

func getFireForecast(place string) ([]bool, error) {
	runtime.GOMAXPROCS(24)
	// Timing code:
	// timeParallelStart := time.Now()
	// for _, place := range places {
	// 	_ = getExpected(place, 1992, 2015, true)
	// }
	// timeParallelEnd := time.Now()
	// fmt.Println("Parallel done")
	// timeSeqStart := time.Now()
	// for _, place := range places {
	// 	_ = getExpected(place, 1992, 2015, false)
	// }
	// timeSeqEnd := time.Now()
	// timeParallel := timeParallelEnd.Sub(timeParallelStart)
	// timeSeq := timeSeqEnd.Sub(timeSeqStart)
	// fmt.Println("Parallel took: ", timeParallel)
	// fmt.Println("Seq took: ", timeSeq)
	// place := "placeholder lmao"
	observedFires, first, last := getObserved(place, 1992, 2015, true)
	if first.IsZero() {
		return nil, errors.New("No data found for given county")
	}
	features := RandomHistoricalWeather(first, last)
	model := linear.NewLogistic(base.BatchGA, 1e-4, 6, 800, features, observedFires)
	model.Learn()
	weathers := Forecast(place)
	fmt.Println(len(weathers))
	forecasts := make([]bool, 0, 5)
	for i := 0; i < 40; i++ {
		fmt.Println(i)
		j := i
		dayPrediction := 0
		for ; j < i+8; j++ {
			prediction, _ := model.Predict(weathers[i])
			dayPrediction = dayPrediction | int(prediction[0])
		}
		i += 7
		//fmt.Println("Hi")
		pred := dayPrediction != 0
		forecasts = append(forecasts, pred)

	}
	fmt.Println(forecasts)
	return forecasts, nil
}
