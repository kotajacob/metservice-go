package metservice

import (
	"context"
	"fmt"
	"net/http"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestGetPollen_Marshal(t *testing.T) {
	testJSONMarshal(t, &Pollen{}, "{}")

	u := &Pollen{
		Location: String("a"),
		PollenDays: []PollenDay{
			{
				DayDescriptor: String("a"),
				Level:         String("b"),
				Type:          String("c"),
				ValidFrom:     &Timestamp{referenceTime},
				ValidTo:       &Timestamp{referenceTime},
			},
		},
		Enabled: Bool(true),
	}

	want := `{
		"location": "a",
		"pollen": [
			{
				"dayDescriptor": "a",
				"level": "b",
				"type": "c",
				"validFromISO": ` + referenceTimeStr + `,
				"validToISO": ` + referenceTimeStr + `
			}
		],
		"pollenEnabled": true
}`

	testJSONMarshal(t, u, want)
}

func TestGetPollen_Get(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	mux.HandleFunc("/pollen_town_Dunedin", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, `{"location": "Dunedin"}`)
	})

	ctx := context.Background()
	pollen, _, err := client.GetPollen(ctx, "Dunedin")
	if err != nil {
		t.Errorf("Client.GetPollen returned error: %v", err)
	}

	want := &Pollen{Location: String("Dunedin")}
	if !cmp.Equal(pollen, want) {
		t.Errorf("Client.GetPollen returned %+v, want %+v", pollen, want)
	}
}
