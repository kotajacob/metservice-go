package metservice

import (
	"context"
	"fmt"
	"net/http"
)

// Forecast represents a metservice forecast.
type Forecast struct {
	// TODO: Implement targeting field
	Days                 []ForecastDay `json:"days,omitempty"`
	LocationGFS          *int          `json:"locationGFS,string,omitempty"`
	LocationIPS          *string       `json:"locationIPS,omitempty"`
	LocationWASP         *string       `json:"locationWASP,omitempty"`
	SaturdayForecastWord *string       `json:"saturdayForecastWord,omitempty"`
	SundayForcastWord    *string       `json:"sundayForecastWord,omitempty"`
}

// ForecastDay represents a day in a Forecast.
type ForecastDay struct {
	Date         *Timestamp `json:"dateISO,omitempty"`
	Forecast     *string    `json:"forecast,omitempty"`
	ForecastWord *string    `json:"forecastWord,omitempty"`
	IssuedAt     *Timestamp `json:"issuedAtISO,omitempty"`
	Max          *int       `json:"max,string,omitempty"`
	Min          *int       `json:"min,string,omitempty"`
	Part         *DayPart   `json:"partDayData,omitempty"`
	RiseSet      *RiseSet   `json:"riseSet,omitempty"`
	Source       *string    `json:"source,omitempty"`
	SourceTemps  *string    `json:"sourceTemps,omitempty"`
}

// ForecastHour represents forecast data for a specific hour. This data
// is obtained from GetObservationForecastHours.
type ForecastHour struct {
	Date          *Timestamp `json:"dateISO,omitempty"`
	Humidity      *int       `json:"humidity,string,omitempty"`
	Offset        *int       `json:"offset,omitempty"`
	Rainfall      *float64   `json:"rainfall,string,omitempty"`
	Temp          *int       `json:"temperature,string,omitempty"`
	WindDirection *string    `json:"windDirection,omitempty"`
	WindSpeed     *int       `json:"windSpeed,string,omitempty"`
}

// DayPart contains DayPartTimes for parts of a ForecastDay.
type DayPart struct {
	Afternoon *DayPartTime `json:"afternoon,omitempty"`
	Evening   *DayPartTime `json:"evening,omitempty"`
	Morning   *DayPartTime `json:"morning,omitempty"`
	Overnight *DayPartTime `json:"overnight,omitempty"`
}

// DayPartTime contains a forecast word and icon type.
type DayPartTime struct {
	ForecastWord *string `json:"forecastWord,omitempty"`
	IconType     *string `json:"iconType,omitempty"`
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
