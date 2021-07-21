package metservice

import (
	"context"
	"fmt"
	"net/http"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestForecast_Marshal(t *testing.T) {
	testJSONMarshal(t, &Forecast{}, "{}")

	u := &Forecast{
		Days: []ForecastDay{
			{
				Date:         &Timestamp{referenceTime},
				Forecast:     String("aa"),
				ForecastWord: String("bb"),
				IssuedAt:     &Timestamp{referenceTime},
				Max:          Int(2),
				Min:          Int(1),
				Part: &DayPart{
					Afternoon: &DayPartTime{
						ForecastWord: String("aaa"),
						IconType:     String("bbb"),
					},
					Evening: &DayPartTime{
						ForecastWord: String("aaa"),
						IconType:     String("bbb"),
					},
					Morning: &DayPartTime{
						ForecastWord: String("aaa"),
						IconType:     String("bbb"),
					},
					Overnight: &DayPartTime{
						ForecastWord: String("aaa"),
						IconType:     String("bbb"),
					},
				},
				RiseSet: &DayRiseSet{
					Date:       &Timestamp{referenceTime},
					FirstLight: &Timestamp{referenceTime},
					ID:         String("aaa"),
					LastLight:  &Timestamp{referenceTime},
					Location:   String("bbb"),
					MoonRise:   &Timestamp{referenceTime},
					MoonSet:    &Timestamp{referenceTime},
					SunRise:    &Timestamp{referenceTime},
					SunSet:     &Timestamp{referenceTime},
				},
				Source:      String("cc"),
				SourceTemps: String("dd"),
			},
		},
		LocationGFS:          Int(123),
		LocationIPS:          String("a"),
		LocationWASP:         String("b"),
		SaturdayForecastWord: String("c"),
		SundayForcastWord:    String("d"),
	}

	want := `{
	"days": [
		{
			"dateISO": ` + referenceTimeStr + `,
			"forecast": "aa",
			"forecastWord": "bb",
			"issuedAtISO": ` + referenceTimeStr + `,
			"max": "2",
			"min": "1",
			"partDayData": {
				"afternoon": {
					"forecastWord": "aaa",
					"iconType": "bbb"
				},
				"evening": {
					"forecastWord": "aaa",
					"iconType": "bbb"
				},
				"morning": {
					"forecastWord": "aaa",
					"iconType": "bbb"
				},
				"overnight": {
					"forecastWord": "aaa",
					"iconType": "bbb"
				}
			},
			"riseSet": {
				"dayISO": ` + referenceTimeStr + `,
				"firstLightISO": ` + referenceTimeStr + `,
				"ID": "aaa",
				"lastLightISO": ` + referenceTimeStr + `,
				"location": "bbb",
				"MoonRiseISO": ` + referenceTimeStr + `,
				"MoonSetISO": ` + referenceTimeStr + `,
				"SunRiseISO": ` + referenceTimeStr + `,
				"SunSetISO": ` + referenceTimeStr + `
			},
			"source": "cc",
			"sourceTemps": "dd"
		}
	],
	"locationGFS": "123",
	"locationIPS": "a",
	"locationWASP": "b",
	"saturdayForecastWord": "c",
	"sundayForecastWord": "d"
}`

	testJSONMarshal(t, u, want)
}

func TestForecast_Get(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	mux.HandleFunc("/localForecastDunedin", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, `{"locationIPS": "DUNEDIN"}`)
	})

	ctx := context.Background()
	forecast, _, err := client.Forecast(ctx, "Dunedin")
	if err != nil {
		t.Errorf("Client.Forecast returned error: %v", err)
	}

	want := &Forecast{LocationIPS: String("DUNEDIN")}
	if !cmp.Equal(forecast, want) {
		t.Errorf("Client.Forecast returned %+v, want %+v", forecast, want)
	}
}
