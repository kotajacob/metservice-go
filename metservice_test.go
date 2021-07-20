package metservice

import (
	"context"
	"flag"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
	"time"

	"github.com/google/go-cmp/cmp"
)

var (
	loc, _          = time.LoadLocation("NZ") // ignore error for testing
	testDunedinJSON = []byte(`{"days":[{"date":"16 Jul","dateISO":"2021-07-16T12:00:00+12:00","dow":"Friday","dowTLA":"Fri","forecast":"Mostly cloudy. A few spots of rain possible again this evening. Northeasterlies, strong about the coast, easing this evening.","forecastWord":"Partly cloudy","issuedAt":"9:09am 16 Jul","issuedAtISO":"2021-07-16T09:09:00+12:00","issuedAtRaw":1626383340000,"max":"13","min":"7","partDayData":{"afternoon":{"forecastWord":"Cloudy","iconType":"DAY"},"evening":{"forecastWord":"Cloudy","iconType":"NIGHT"},"morning":{"forecastWord":"Partly cloudy","iconType":"DAY"},"overnight":{"forecastWord":"Showers","iconType":"NIGHT"}},"riseSet":{"day":"16 July 2021","dayISO":"2021-07-16T00:00:00+12:00","firstLight":"7:39am","firstLightISO":"2021-07-16T07:39:00+12:00","id":"loc1036728","lastLight":"5:50pm","lastLightISO":"2021-07-16T17:50:00+12:00","location":"Dunedin","moonRise":"11:37am","moonRiseISO":"2021-07-16T11:37:00+12:00","moonSet":"11:55pm","moonSetISO":"2021-07-16T23:55:00+12:00","sunRise":"8:12am","sunRiseHour":8,"sunRiseISO":"2021-07-16T08:12:00+12:00","sunSet":"5:16pm","sunSetHour":17,"sunSetISO":"2021-07-16T17:16:00+12:00"},"source":"IPS","sourceTemps":"IPS"},{"date":"17 Jul","dateISO":"2021-07-17T12:00:00+12:00","dow":"Saturday","dowTLA":"Sat","forecast":"Low cloud. Occasional rain developing early morning. Winds mainly light.","forecastWord":"Showers","issuedAt":"11:13am 16 Jul","issuedAtISO":"2021-07-16T11:13:00+12:00","issuedAtRaw":1626390780000,"max":"11","min":"8","partDayData":{"afternoon":{"forecastWord":"Showers","iconType":"DAY"},"evening":{"forecastWord":"Showers","iconType":"NIGHT"},"morning":{"forecastWord":"Showers","iconType":"DAY"},"overnight":{"forecastWord":"Few showers","iconType":"NIGHT"}},"riseSet":{"day":"17 July 2021","dayISO":"2021-07-17T00:00:00+12:00","firstLight":"7:39am","firstLightISO":"2021-07-17T07:39:00+12:00","id":"loc1036728","lastLight":"5:50pm","lastLightISO":"2021-07-17T17:50:00+12:00","location":"Dunedin","moonRise":"11:59am","moonRiseISO":"2021-07-17T11:59:00+12:00","sunRise":"8:12am","sunRiseHour":8,"sunRiseISO":"2021-07-17T08:12:00+12:00","sunSet":"5:17pm","sunSetHour":17,"sunSetISO":"2021-07-17T17:17:00+12:00"},"source":"IPS","sourceTemps":"IPS"}],"locationGFS":"93892","locationIPS":"DUNEDIN","locationWASP":"DUNEDIN","saturdayForecastWord":"Showers","sundayForecastWord":"Few showers"}`)
	testDunedin     = Forecast{
		Days: []Day{
			{
				Date:         time.Date(2021, time.July, 16, 12, 0, 0, 0, loc),
				Forecast:     "Mostly cloudy. A few spots of rain possible again this evening. Northeasterlies, strong about the coast, easing this evening.",
				ForecastWord: "Partly cloudy",
				IssuedAt:     time.Date(2021, time.July, 16, 9, 9, 0, 0, loc),
				Max:          13,
				Min:          7,
				Source:       "IPS",
				SourceTemps:  "IPS",
			},
			{
				Date:         time.Date(2021, time.July, 17, 12, 0, 0, 0, loc),
				Forecast:     "Low cloud. Occasional rain developing early morning. Winds mainly light.",
				ForecastWord: "Showers",
				IssuedAt:     time.Date(2021, time.July, 16, 11, 13, 0, 0, loc),
				Max:          11,
				Min:          8,
				Source:       "IPS",
				SourceTemps:  "IPS",
			},
		},
		LocationGFS:          93892,
		LocationIPS:          "DUNEDIN",
		LocationWASP:         "DUNEDIN",
		SaturdayForecastWord: "Showers",
		SundayForcastWord:    "Few showers",
	}
)

func localServer() *httptest.Server {
	writeJson := func(b []byte) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
			w.Write(b)
		})
	}

	mux := http.NewServeMux()
	mux.Handle("/localForecastDunedin", writeJson(testDunedinJSON))

	return httptest.NewUnstartedServer(mux)
}

var localClient *Client

func TestMain(m *testing.M) {
	flag.Parse()

	svr := localServer()
	svr.Start()
	defer svr.Close()

	localClient = &Client{
		HTTPClient: http.DefaultClient,
		BaseURL:    svr.URL,
	}

	os.Exit(m.Run())
}

func TestForecast(t *testing.T) {
	forecast, err := localClient.Forecast(context.Background(), "Dunedin")
	if err != nil {
		t.Errorf("%v", err)
		return
	}
	if diff := cmp.Diff(testDunedin, forecast); diff != "" {
		t.Errorf("forecasts do not match:\n%s", diff)
		return
	}
}
