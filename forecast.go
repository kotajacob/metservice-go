package metservice

import (
	"context"
	"fmt"
	"net/http"
)

// Forecast represents a metservice forecast.
type Forecast struct {
	// TODO: Implement targeting field
	Days                 []ForecastDay `json:"days"`
	LocationGFS          *int          `json:"locationGFS,string"`
	LocationIPS          *string       `json:"locationIPS"`
	LocationWASP         *string       `json:"locationWASP"`
	SaturdayForecastWord *string       `json:"saturdayForecastWord"`
	SundayForcastWord    *string       `json:"sundayForecastWord"`
}

// ForecastDay represents a day in a Forecast.
type ForecastDay struct {
	Date         *Timestamp `json:"dateISO"`
	Forecast     *string    `json:"forecast"`
	ForecastWord *string    `json:"forecastWord"`
	IssuedAt     *Timestamp `json:"issuedAtISO"`
	Max          *int       `json:"max,string"`
	Min          *int       `json:"min,string"`
	Part         *DayPart   `json:"partDayData"`
	RiseSet      *RiseSet   `json:"riseSet"`
	Source       *string    `json:"source"`
	SourceTemps  *string    `json:"sourceTemps"`
}

// ForecastHour represents forecast data for a specific hour. This data
// is obtained from GetObservationForecastHours.
type ForecastHour struct {
	Date          *Timestamp `json:"dateISO"`
	Humidity      *int       `json:"humidity,string"`
	Offset        *int       `json:"offset"`
	Rainfall      *float64   `json:"rainfall,string"`
	Temp          *int       `json:"temperature,string"`
	WindDirection *string    `json:"windDirection"`
	WindSpeed     *int       `json:"windSpeed,string"`
}

// DayPart contains DayPartTimes for parts of a ForecastDay.
type DayPart struct {
	Afternoon *DayPartTime `json:"afternoon"`
	Evening   *DayPartTime `json:"evening"`
	Morning   *DayPartTime `json:"morning"`
	Overnight *DayPartTime `json:"overnight"`
}

// DayPartTime contains a forecast word and icon type.
type DayPartTime struct {
	ForecastWord *string `json:"forecastWord"`
	IconType     *string `json:"iconType"`
}

// GetForecast gets a Forecast for a given location using a context. The location
// string should be capitalized - i.e. Dunedin. A list of possible locations
// can be found here https://www.metservice.com/towns-cities/
func (c *Client) GetForecast(ctx context.Context, location string) (*Forecast, *http.Response, error) {
	forecast := new(Forecast)
	path := fmt.Sprintf("localForecast%s", location)
	rsp, err := c.Do(ctx, path, forecast)
	if err != nil {
		return &Forecast{}, rsp, err
	}
	return forecast, rsp, nil
}
