package metservice

import (
	"context"
	"fmt"
	"net/http"
)

// RiseSet represents the sun/moon rise and set times for a day.
type RiseSet struct {
	Date       *Timestamp `json:"dayISO"`
	FirstLight *Timestamp `json:"firstLightISO"`
	ID         *string    `json:"id"`
	LastLight  *Timestamp `json:"lastLightISO"`
	Location   *string    `json:"location"`
	MoonRise   *Timestamp `json:"moonRiseISO"`
	MoonSet    *Timestamp `json:"moonSetISO"`
	SunRise    *Timestamp `json:"sunRiseISO"`
	SunSet     *Timestamp `json:"sunSetISO"`
}

// GetRiseSet gets a RiseSet representing the sun/moon rise and set times for
// the current day for a given location. The location string should be
// capitalized - i.e. Dunedin. A list of possible locations can be found here
// https://www.metservice.com/towns-cities/
func (c *Client) GetRiseSet(ctx context.Context, location string) (*RiseSet, *http.Response, error) {
	riseSet := new(RiseSet)
	path := fmt.Sprintf("riseSet_%s", location)
	rsp, err := c.Do(ctx, path, riseSet)
	if err != nil {
		return &RiseSet{}, rsp, err
	}
	return riseSet, rsp, nil
}
