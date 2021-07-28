package metservice

import (
	"context"
	"fmt"
	"net/http"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestGetObservation_Marshal(t *testing.T) {
	testJSONMarshal(t, &Observation{}, "{}")

	u := &Observation{
		ID:         String("a"),
		Location:   String("b"),
		LocationID: Int(1),
		ThreeHour: &ObservationThreeHour{
			ClothingLayers:  Int(11),
			Date:            &Timestamp{referenceTime},
			Humidity:        Int(22),
			Pressure:        String("aa"),
			Rainfall:        Float64(3.3),
			Temp:            Int(44),
			WindChill:       Int(55),
			WindDirection:   String("bb"),
			WindProofLayers: Int(66),
			WindSpeed:       Int(77),
		},
		TwentyFourHour: &ObservationTwentyFourHour{
			DatePretty: String("aa"),
			Max:        Int(11),
			Min:        Int(22),
			Rainfall:   Float64(3.3),
		},
	}

	want := `{
	"id": "a",
	"location": "b",
	"locationId": "1",
	"threeHour": {
		"clothingLayers": "11",
		"dateTimeISO": ` + referenceTimeStr + `,
		"humidity": "22",
		"pressure": "aa",
		"rainfall": "3.3",
		"temp": "44",
		"windChill": "55",
		"windDirection": "bb",
		"windProofLayers": "66",
		"windSpeed": "77"
	},
	"twentyFourHour" : {
		"dateTime": "aa",
		"maxTemp": 11,
		"minTemp": 22,
		"rainfall": "3.3"
	}
}`

	testJSONMarshal(t, u, want)
}

func TestGetObservationOneMin_Marshal(t *testing.T) {
	testJSONMarshal(t, &ObservationOneMin{}, "{}")

	u := &ObservationOneMin{
		ClothingLayers:   Int(1),
		Current:          Bool(true),
		Past:             String("a"),
		Rainfall:         Float64(2.2),
		RelativeHumidity: Int(3),
		Status:           String("b"),
		Date:             &Timestamp{referenceTime},
		WindProofLayers:  Int(4),
	}

	want := `{
	"clothingLayers": "1",
	"isObservationCurrent": true,
	"past": "a",
	"rainfall": "2.2",
	"relativeHumidity": "3",
	"status": "b",
	"timeISO": ` + referenceTimeStr + `,
	"windProofLayers": "4"
}`

	testJSONMarshal(t, u, want)
}

func TestGetObservationForecastHours_Marshal(t *testing.T) {
	testJSONMarshal(t, &ObservationForecastHours{}, "{}")

	u := &ObservationForecastHours{
		Observations: []ObservationHour{
			{
				Date:          &Timestamp{referenceTime},
				Offset:        Int(1),
				Rainfall:      Float64(2.2),
				Temp:          Float64(3.3),
				WindDirection: String("a"),
				WindSpeed:     Int(4),
			},
		},
		Forecasts: []ForecastHour{
			{
				Date:          &Timestamp{referenceTime},
				Humidity:      Int(1),
				Offset:        Int(2),
				Rainfall:      Float64(3.3),
				Temp:          Int(4),
				WindDirection: String("a"),
				WindSpeed:     Int(4),
			},
		},
		Count:                 Int(1),
		WindSpeed:             Int(2),
		Location:              String("a"),
		LocationName:          String("b"),
		RainfallTotalForecast: Float64(3.3),
		RainfallTotalObserved: Float64(4.4),
	}

	want := `{
	"actualData": [
		{
			"dateISO": ` + referenceTimeStr + `,
			"offset": 1,
			"rainfall": "2.2",
			"temperature": "3.3",
			"windDirection": "a",
			"windSpeed": "4"
		}
	],
	"forecastData": [
		{
			"dateISO": ` + referenceTimeStr + `,
			"humidity": "1",
			"offset": 2,
			"rainfall": "3.3",
			"temperature": "4",
			"windDirection": "a",
			"windSpeed": "4"
		}
	],
	"dataPointCount": 1,
	"latestObsWindSpeed": 2,
	"location": "a",
	"locationName": "b",
	"rainfallTotalForecast": "3.3",
	"rainfallTotalObserved": "4.4"
}`

	testJSONMarshal(t, u, want)
}

func TestGetObservation_Get(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	mux.HandleFunc("/localObs_Dunedin", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, `{"id": "DUNEDIN"}`)
	})

	ctx := context.Background()
	observation, _, err := client.GetObservation(ctx, "Dunedin")
	if err != nil {
		t.Errorf("Client.GetObservation returned error: %v", err)
	}

	want := &Observation{ID: String("DUNEDIN")}
	if !cmp.Equal(observation, want) {
		t.Errorf("Client.GetObservation returned %+v, want %+v", observation, want)
	}
}

func TestGetObservationOneMin_Get(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	mux.HandleFunc("/oneMinObs_Dunedin", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, `{"status": "DUNEDIN"}`)
	})

	ctx := context.Background()
	observation, _, err := client.GetObservationOneMin(ctx, "Dunedin")
	if err != nil {
		t.Errorf("Client.GetObservationOneMin returned error: %v", err)
	}

	want := &ObservationOneMin{Status: String("DUNEDIN")}
	if !cmp.Equal(observation, want) {
		t.Errorf("Client.GetObservationOneMin returned %+v, want %+v", observation, want)
	}
}

func TestGetObservationForecastHours_Get(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	mux.HandleFunc("/hourlyObsAndForecast_Dunedin", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, `{"location": "Dunedin"}`)
	})

	ctx := context.Background()
	observation, _, err := client.GetObservationForecastHours(ctx, "Dunedin")
	if err != nil {
		t.Errorf("Client.GetObservationForecastHours returned error: %v", err)
	}

	want := &ObservationForecastHours{Location: String("Dunedin")}
	if !cmp.Equal(observation, want) {
		t.Errorf("Client.GetObservationForecastHours returned %+v, want %+v", observation, want)
	}
}
