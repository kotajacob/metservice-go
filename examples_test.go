package metservice

import (
	"context"
	"fmt"
	"testing"
)

func ExampleForecast() {
	client := NewClient()
	ctx := context.Background()

	forecast, _, err := client.GetForecast(ctx, "Dunedin")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(*forecast.LocationIPS)
	for _, day := range forecast.Days {
		fmt.Printf("%v\nforecast: %v\nmax: %vC\nmin: %vC\n\n",
			*day.Date,
			*day.ForecastWord,
			*day.Max,
			*day.Min)
	}
}

func TestExampleForecast(t *testing.T) {
	if !testing.Verbose() {
		return
	}
	ExampleForecast()
}
