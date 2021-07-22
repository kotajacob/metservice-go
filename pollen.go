package metservice

import (
	"context"
	"fmt"
	"net/http"
)

// Pollen represents the pollen/alergy data for the next few days.
type Pollen struct {
	Location   *string     `json:"location,omitempty"`
	PollenDays []PollenDay `json:"pollen,omitempty"`
	Enabled    *bool       `json:"pollenEnabled,omitempty"`
}

// PollenDay represents the pollen data for a single day.
type PollenDay struct {
	DayDescriptor *string    `json:"dayDescriptor,omitempty"`
	Level         *string    `json:"level,omitempty"`
	Type          *string    `json:"type,omitempty"`
	ValidFrom     *Timestamp `json:"validFromISO,omitempty"`
	ValidTo       *Timestamp `json:"validToISO,omitempty"`
}

// GetPollen gets a Pollen representing the pollen/alergy data for the next few
// days for a given location.
// The location string should be capitalized - i.e. Dunedin. A list of possible
// locations can be found here https://www.metservice.com/towns-cities/
func (c *Client) GetPollen(ctx context.Context, location string) (*Pollen, *http.Response, error) {
	pollen := new(Pollen)
	path := fmt.Sprintf("pollen_town_%s", location)
	rsp, err := c.Do(ctx, path, pollen)
	if err != nil {
		return &Pollen{}, rsp, err
	}
	return pollen, rsp, nil
}
