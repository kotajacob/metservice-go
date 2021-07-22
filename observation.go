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
	ClothingLayers  *int       `json:"clothingLayers,string,omitempty"`
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
	Max        *int     `json:"maxTemp,string,omitempty"`
	Min        *int     `json:"minTemp,string,omitempty"`
	Rainfall   *float64 `json:"rainfall,string,omitempty"`
}

// GetObservation gets an Observation for a given location. The location string
// should be capitalized - i.e. Dunedin. A list of possible locations can be
// found here https://www.metservice.com/towns-cities/
func (c *Client) GetObservation(ctx context.Context, location string) (*Observation, *http.Response, error) {
	observation := new(Observation)
	path := fmt.Sprintf("localObs_%s", location)
	rsp, err := c.Do(ctx, path, observation)
	if err != nil {
		return &Observation{}, rsp, err
	}
	return observation, rsp, nil
}
