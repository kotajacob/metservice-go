package metservice

import (
	"context"
	"fmt"
	"net/http"
)

// Observation represents a recent metservice observation.
type Observation struct {
	ID             *string                    `json:"id"`
	Location       *string                    `json:"location"`
	LocationID     *int                       `json:"locationId,string"`
	ThreeHour      *ObservationThreeHour      `json:"threeHour"`
	TwentyFourHour *ObservationTwentyFourHour `json:"twentyFourHour"`
}

// ObservationThreeHour represents observation data updated every 3 hours.
type ObservationThreeHour struct {
	ClothingLayers  *string    `json:"clothingLayers"`
	Date            *Timestamp `json:"dateTimeISO"`
	Humidity        *int       `json:"humidity,string"`
	Pressure        *string    `json:"pressure"`
	Rainfall        *float64   `json:"rainfall,string"`
	Temp            *int       `json:"temp,string"`
	WindChill       *int       `json:"windChill,string"`
	WindDirection   *string    `json:"windDirection"`
	WindProofLayers *int       `json:"windProofLayers,string"`
	WindSpeed       *int       `json:"windSpeed,string"`
}

// ObservationTwentyFourHour represents observation data updated day.
type ObservationTwentyFourHour struct {
	DatePretty *string  `json:"dateTime"`
	Max        *int     `json:"maxTemp"`
	Min        *int     `json:"minTemp"`
	Rainfall   *float64 `json:"rainfall,string"`
}

// ObservationForecastHours represents observation and forecast data hourly,
// usually for around 48 hours with 9 or 10 observations and the rest being
// forecasts. I felt some fields were redundant so I ignored them.
type ObservationForecastHours struct {
	Observations          []ObservationHour `json:"actualData"`
	Forecasts             []ForecastHour    `json:"forecastData"`
	Count                 *int              `json:"dataPointCount"`
	WindSpeed             *int              `json:"latestObsWindSpeed"`
	Location              *string           `json:"location"`
	LocationName          *string           `json:"locationName"`
	RainfallTotalForecast *float64          `json:"rainfallTotalForecast"`
	RainfallTotalObserved *float64          `json:"rainfallTotalObserved"`
}

// ObservationHour represents observation data for a specific hour. This data
// is obtained from GetObservationForecastHours.
type ObservationHour struct {
	Date          *Timestamp `json:"dateISO"`
	Offset        *int       `json:"offset"`
	Rainfall      *float64   `json:"rainfall,string"`
	Temp          *float64   `json:"temperature,string"`
	WindDirection *string    `json:"windDirection"`
	WindSpeed     *int       `json:"windSpeed,string"`
}

// ObservationOneMin represents observation data updated to the minute. It has
// less detail than the daily observations.
type ObservationOneMin struct {
	ClothingLayers   *string    `json:"clothingLayers"`
	Current          *bool      `json:"isObservationCurrent"`
	Past             *string    `json:"past"`
	Rainfall         *float64   `json:"rainfall,string"`
	RelativeHumidity *int       `json:"relativeHumidity,string"`
	Status           *string    `json:"status"`
	Date             *Timestamp `json:"timeISO"`
	WindProofLayers  *int       `json:"windProofLayers,string"`
}

// GetObservation gets an Observation for a given location.
// The location string should be capitalized - i.e. Dunedin. A list of possible
// locations can be found here https://www.metservice.com/towns-cities/
func (c *Client) GetObservation(ctx context.Context, location string) (*Observation, *http.Response, error) {
	observation := new(Observation)
	path := fmt.Sprintf("localObs_%s", location)
	rsp, err := c.Do(ctx, path, observation)
	if err != nil {
		return &Observation{}, rsp, err
	}
	return observation, rsp, nil
}

// GetObservationForecastHours gets an ObservationForecastHours containing
// hourly observations and forecasts for about a 48 hour period for a specific
// location.
func (c *Client) GetObservationForecastHours(ctx context.Context, location string) (*ObservationForecastHours, *http.Response, error) {
	ofh := new(ObservationForecastHours)
	path := fmt.Sprintf("hourlyObsAndForecast_%s", location)
	rsp, err := c.Do(ctx, path, ofh)
	if err != nil {
		return &ObservationForecastHours{}, rsp, err
	}
	return ofh, rsp, nil
}

// GetObservationOneMin gets an Observation for a given location.
// The location string should be capitalized - i.e. Dunedin. A list of possible
// locations can be found here https://www.metservice.com/towns-cities/
func (c *Client) GetObservationOneMin(ctx context.Context, location string) (*ObservationOneMin, *http.Response, error) {
	observation := new(ObservationOneMin)
	path := fmt.Sprintf("oneMinObs_%s", location)
	rsp, err := c.Do(ctx, path, observation)
	if err != nil {
		return &ObservationOneMin{}, rsp, err
	}
	return observation, rsp, nil
}
