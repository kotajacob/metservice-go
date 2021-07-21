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

// ForecastDay represents a day in a forecast.
type ForecastDay struct {
	// TODO: Implement partDayData
	Date         *Timestamp  `json:"dateISO,omitempty"`
	Forecast     *string     `json:"forecast,omitempty"`
	ForecastWord *string     `json:"forecastWord,omitempty"`
	IssuedAt     *Timestamp  `json:"issuedAtISO,omitempty"`
	Max          *int        `json:"max,string,omitempty"`
	Min          *int        `json:"min,string,omitempty"`
	RiseSet      *DayRiseSet `json:"riseSet,omitempty"`
	Source       *string     `json:"source,omitempty"`
	SourceTemps  *string     `json:"sourceTemps,omitempty"`
}

// Rises represents the sun/moon rise and set times in a day.
type DayRiseSet struct {
	Date       *Timestamp `json:"dayISO,omitempty"`
	FirstLight *Timestamp `json:"firstLightISO,omitempty"`
	ID         *string    `json:"id,omitempty"`
	LastLight  *Timestamp `json:"lastLightISO,omitempty"`
	Location   *string    `json:"location,omitempty"`
	MoonRise   *Timestamp `json:"moonRiseISO,omitempty"`
	MoonSet    *Timestamp `json:"moonSetISO,omitempty"`
	SunRise    *Timestamp `json:"sunRiseISO,omitempty"`
	SunSet     *Timestamp `json:"sunSetISO,omitempty"`
}

// Forecast gets a Forecast for a given location using a context. The location
// string should be capitalized - i.e. Dunedin. A list of possible locations
// can be found here https://www.metservice.com/towns-cities/
func (c *Client) Forecast(ctx context.Context, location string) (*Forecast, *http.Response, error) {
	forecast := new(Forecast)
	path := fmt.Sprintf("localForecast%s", location)
	rsp, err := c.Do(ctx, path, forecast)
	if err != nil {
		return &Forecast{}, rsp, err
	}

	return forecast, rsp, nil
}
