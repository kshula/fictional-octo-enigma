package main

import (
	"fmt"
	"sort"
	"time"
)

// DataPoint represents a single data point in the time series.
type DataPoint struct {
	Timestamp time.Time
	Value     float64
}

// TimeSeries represents the time series data.
type TimeSeries struct {
	Data []DataPoint
}

// Insert adds a new data point to the time series.
func (ts *TimeSeries) Insert(timestamp time.Time, value float64) {
	dataPoint := DataPoint{Timestamp: timestamp, Value: value}
	ts.Data = append(ts.Data, dataPoint)
	// Ensure data is sorted by timestamp
	sort.Slice(ts.Data, func(i, j int) bool {
		return ts.Data[i].Timestamp.Before(ts.Data[j].Timestamp)
	})
}

// Query retrieves data points within the specified time range.
func (ts *TimeSeries) Query(start, end time.Time) []DataPoint {
	var result []DataPoint
	for _, dataPoint := range ts.Data {
		if dataPoint.Timestamp.After(start) && dataPoint.Timestamp.Before(end) {
			result = append(result, dataPoint)
		}
	}
	return result
}

func main() {
	// Create a new time series
	ts := &TimeSeries{}

	// Insert sample data
	ts.Insert(time.Now().Add(-10*time.Minute), 1.1)
	ts.Insert(time.Now().Add(-8*time.Minute), 1.2)
	ts.Insert(time.Now().Add(-5*time.Minute), 1.3)
	ts.Insert(time.Now().Add(-3*time.Minute), 1.4)
	ts.Insert(time.Now(), 1.5)

	// Define query range
	start := time.Now().Add(-9 * time.Minute)
	end := time.Now().Add(-4 * time.Minute)

	// Query data
	result := ts.Query(start, end)

	// Print query result
	fmt.Println("Query Results:")
	for _, dp := range result {
		fmt.Printf("Timestamp: %s, Value: %.2f\n", dp.Timestamp.Format(time.RFC3339), dp.Value)
	}
}
