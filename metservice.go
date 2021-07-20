package metservice

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"time"
)

// BaseURL is the default base URL for the metservice JSON API.
const BaseURL = "https://www.metservice.com/publicData"

// Client used for metservice-go.
type Client struct {
	HTTPClient *http.Client
	BaseURL    string
}

// NewClient constructs a client using http.DefaultClient and the default
// base URL. The returned client is ready for use.
func NewClient() *Client {
	return &Client{
		HTTPClient: http.DefaultClient,
		BaseURL:    BaseURL,
	}
}

type forecastResponce struct {
	// TODO: Implement targeting field
	Days                 []dayResponce `json:"days"`
	LocationGFS          string        `json:"locationGFS"`
	LocationIPS          string        `json:"locationIPS"`
	LocationWASP         string        `json:"locationWASP"`
	SaturdayForecastWord string        `json:"saturdayForecastWord"`
	SundayForcastWord    string        `json:"sundayForecastWord"`
}

// some of the time fields are ignored in place of the ISO field.
type dayResponce struct {
	// TODO: Implement partDayData and riseSet
	DateISO      string `json:"dateISO"`
	Forecast     string `json:"forecast"`
	ForecastWord string `json:"forecastWord"`
	IssuedAtISO  string `json:"issuedAtISO"`
	Max          string `json:"max"`
	Min          string `json:"min"`
	Source       string `json:"source"`
	SourceTemps  string `json:"sourceTemps"`
}

// Forecast contains information of a specific weather forcast.
type Forecast struct {
	Days                 []Day
	LocationGFS          int
	LocationIPS          string
	LocationWASP         string
	SaturdayForecastWord string
	SundayForcastWord    string
}

// Day represents forcast information about a day.
type Day struct {
	Date         time.Time
	Forecast     string
	ForecastWord string
	IssuedAt     time.Time
	Max          int
	Min          int
	Source       string
	SourceTemps  string
}

// StatusError is returned when a bad responce code is received from the API.
type StatusError struct {
	Code int
}

var _ error = StatusError{}

func (e StatusError) Error() string {
	return fmt.Sprintf("bad responce status code: %d", e.Code)
}

// GetForecast returns a Forecast for a given location using a context. The
// location string should be capitalized - i.e. Dunedin.
func (c *Client) GetForecast(ctx context.Context, location string) (Forecast, error) {
	reqPath := fmt.Sprintf("/localForecast%s", location)
	req, err := http.NewRequest("GET", c.BaseURL+reqPath, nil)
	if err != nil {
		return Forecast{}, fmt.Errorf("failed to build request: %v", err)
	}
	req = req.WithContext(ctx)

	rsp, err := c.HTTPClient.Do(req)
	if err != nil {
		return Forecast{}, fmt.Errorf("failed to do request: %v", err)
	}
	defer rsp.Body.Close()

	if rsp.StatusCode != http.StatusOK {
		return Forecast{}, StatusError{Code: rsp.StatusCode}
	}

	var fr forecastResponce
	if err := json.NewDecoder(rsp.Body).Decode(&fr); err != nil {
		return Forecast{}, fmt.Errorf("failed to parse json: %v", err)
	}

	// check for errors closing responce
	// https://www.joeshaw.org/dont-defer-close-on-writable-files/
	if err := rsp.Body.Close(); err != nil {
		return Forecast{}, err
	}

	f := Forecast{
		LocationIPS:          fr.LocationIPS,
		LocationWASP:         fr.LocationWASP,
		SaturdayForecastWord: fr.SaturdayForecastWord,
		SundayForcastWord:    fr.SundayForcastWord,
	}
	locationGFS, err := strconv.Atoi(fr.LocationGFS)
	if err != nil {
		return Forecast{}, fmt.Errorf("failed to parse locationGFS: %v", err)
	}
	f.LocationGFS = locationGFS

	// convert non-string values in forecastResponce's days
	for _, v := range fr.Days {
		date, err := time.Parse(time.RFC3339, v.DateISO)
		if err != nil {
			return Forecast{}, fmt.Errorf("failed to parse day dateISO: %v", err)
		}
		issuedAt, err := time.Parse(time.RFC3339, v.IssuedAtISO)
		if err != nil {
			return Forecast{}, fmt.Errorf("failed to parse day issuedAtISO: %v", err)
		}
		max, err := strconv.Atoi(v.Max)
		if err != nil {
			return Forecast{}, fmt.Errorf("failed to parse day max: %v", err)
		}
		min, err := strconv.Atoi(v.Min)
		if err != nil {
			return Forecast{}, fmt.Errorf("failed to parse day min: %v", err)
		}
		f.Days = append(f.Days, Day{
			Date:         date,
			Forecast:     v.Forecast,
			ForecastWord: v.ForecastWord,
			IssuedAt:     issuedAt,
			Max:          max,
			Min:          min,
			Source:       v.Source,
			SourceTemps:  v.SourceTemps,
		})
	}

	return f, nil
}
