package main

import (
	"time"
)

const (
	//AddBufferSize is the size of the buffer for the addDatesChan channel
	AddBufferSize = 32
)

//ParallelResultSlice allows for parallel read of results
type ParallelResultSlice struct {
	dates        []time.Time
	addDatesChan chan []time.Time
	killChan     chan int
}

//Listen for requests to add data to the slice
func (rs ParallelResultSlice) Listen(finalResult chan []time.Time) {
	var (
		dates []time.Time
	)
	for {
		select {
		case dates = <-rs.addDatesChan:
			rs.dates = append(rs.dates, dates...)
		case <-rs.killChan:
			finalResult <- rs.dates
			break
		}
	}
}

//NewParallelResultSlice intializes values for the result slice struct
func NewParallelResultSlice() ParallelResultSlice {
	var resultSlice ParallelResultSlice
	resultSlice.dates = make([]time.Time, 0)
	resultSlice.killChan = make(chan int)
	resultSlice.addDatesChan = make(chan []time.Time, AddBufferSize)
	return resultSlice
}

//AddDates adds dates to the central dates slice
func (rs ParallelResultSlice) AddDates(dates []time.Time) {
	rs.addDatesChan <- dates
	return
}

//Kill ends the main loop
func (rs ParallelResultSlice) Kill() {
	rs.killChan <- 1
	return
}

//GetDates returns the underlying slice of dates
func (rs ParallelResultSlice) GetDates() []time.Time {
	return rs.dates
}
