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
		"maxTemp": "11",
		"minTemp": "22",
		"rainfall": "3.3"
	}
}`

	testJSONMarshal(t, u, want)
}

func TestGetObservation_Get(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	mux.HandleFunc("/localObs_Dunedin", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, `{"ID": "DUNEDIN"}`)
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
