package metservice

import (
	"context"
	"fmt"
	"net/http"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestGetRiseSet_Marshal(t *testing.T) {
	testJSONMarshal(t, &RiseSet{}, "{}")

	u := &RiseSet{
		Date:       &Timestamp{referenceTime},
		FirstLight: &Timestamp{referenceTime},
		ID:         String("aaa"),
		LastLight:  &Timestamp{referenceTime},
		Location:   String("bbb"),
		MoonRise:   &Timestamp{referenceTime},
		MoonSet:    &Timestamp{referenceTime},
		SunRise:    &Timestamp{referenceTime},
		SunSet:     &Timestamp{referenceTime},
	}

	want := `{
		"dayISO": ` + referenceTimeStr + `,
		"firstLightISO": ` + referenceTimeStr + `,
		"ID": "aaa",
		"lastLightISO": ` + referenceTimeStr + `,
		"location": "bbb",
		"MoonRiseISO": ` + referenceTimeStr + `,
		"MoonSetISO": ` + referenceTimeStr + `,
		"SunRiseISO": ` + referenceTimeStr + `,
		"SunSetISO": ` + referenceTimeStr + `
}`

	testJSONMarshal(t, u, want)
}

func TestGetRiseSet_Get(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	mux.HandleFunc("/riseSet_Dunedin", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, `{"ID": "TEST"}`)
	})

	ctx := context.Background()
	riseSet, _, err := client.GetRiseSet(ctx, "Dunedin")
	if err != nil {
		t.Errorf("Client.GetRiseSet returned error: %v", err)
	}

	want := &RiseSet{ID: String("TEST")}
	if !cmp.Equal(riseSet, want) {
		t.Errorf("Client.GetRiseSet returned %+v, want %+v", riseSet, want)
	}
}
