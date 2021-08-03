package metservice

import (
	"context"
	"fmt"
	"net/http"
)

// Observation represents a recent metservice observation.
type Observation struct {
	ID             *string                    `json:"id,omitempty"`
	Location       *string                    `json:"location,omitempty"`
	LocationID     *int                       `json:"locationId,string,omitempty"`
	ThreeHour      *ObservationThreeHour      `json:"threeHour,omitempty"`
	TwentyFourHour *ObservationTwentyFourHour `json:"twentyFourHour,omitempty"`
}

// ObservationThreeHour represents observation data updated every 3 hours.
type ObservationThreeHour struct {
	ClothingLayers  *string    `json:"clothingLayers,omitempty"`
	Date            *Timestamp `json:"dateTimeISO,omitempty"`
	Humidity        *int       `json:"humidity,string,omitempty"`
	Pressure        *string    `json:"pressure,omitempty"`
	Rainfall        *float64   `json:"rainfall,string,omitempty"`
	Temp            *int       `json:"temp,string,omitempty"`
	WindChill       *int       `json:"windChill,string,omitempty"`
	WindDirection   *string    `json:"windDirection,omitempty"`
	WindProofLayers *int       `json:"windProofLayers,string,omitempty"`
	WindSpeed       *int       `json:"windSpeed,string,omitempty"`
}

// ObservationTwentyFourHour represents observation data updated day.
type ObservationTwentyFourHour struct {
	DatePretty *string  `json:"dateTime,omitempty"`
	Max        *int     `json:"maxTemp,omitempty"`
	Min        *int     `json:"minTemp,omitempty"`
	Rainfall   *float64 `json:"rainfall,string,omitempty"`
}

// ObservationForecastHours represents observation and forecast data hourly,
// usually for around 48 hours with 9 or 10 observations and the rest being
// forecasts. I felt some fields were redundant so I ignored them.
type ObservationForecastHours struct {
	Observations          []ObservationHour `json:"actualData,omitempty"`
	Forecasts             []ForecastHour    `json:"forecastData,omitempty"`
	Count                 *int              `json:"dataPointCount,omitempty"`
	WindSpeed             *int              `json:"latestObsWindSpeed,omitempty"`
	Location              *string           `json:"location,omitempty"`
	LocationName          *string           `json:"locationName,omitempty"`
	RainfallTotalForecast *float64          `json:"rainfallTotalForecast,string,omitempty"`
	RainfallTotalObserved *float64          `json:"rainfallTotalObserved,string,omitempty"`
}

// ObservationHour represents observation data for a specific hour. This data
// is obtained from GetObservationForecastHours.
type ObservationHour struct {
	Date          *Timestamp `json:"dateISO,omitempty"`
	Offset        *int       `json:"offset,omitempty"`
	Rainfall      *float64   `json:"rainfall,string,omitempty"`
	Temp          *float64   `json:"temperature,string,omitempty"`
	WindDirection *string    `json:"windDirection,omitempty"`
	WindSpeed     *int       `json:"windSpeed,string,omitempty"`
}

// ObservationOneMin represents observation data updated to the minute. It has
// less detail than the daily observations.
type ObservationOneMin struct {
	ClothingLayers   *string    `json:"clothingLayers,omitempty"`
	Current          *bool      `json:"isObservationCurrent,omitempty"`
	Past             *string    `json:"past,omitempty"`
	Rainfall         *float64   `json:"rainfall,string,omitempty"`
	RelativeHumidity *int       `json:"relativeHumidity,string,omitempty"`
	Status           *string    `json:"status,omitempty"`
	Date             *Timestamp `json:"timeISO,omitempty"`
	WindProofLayers  *int       `json:"windProofLayers,string,omitempty"`
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
