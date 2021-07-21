package metservice

import (
	"context"
	"fmt"
)

func ExampleForecast() {
	client := NewClient()
	ctx := context.Background()

	forecast, _, err := client.Forecast(ctx, "Dunedin")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(*forecast.LocationIPS)
}
