package memo

import (
	"ports"
	"reflect"
	"testing"
)

func TestCreateOrUpdate(t *testing.T) {
	testCases := map[string]struct {
		in   *ports.Port
		db   map[string]*ports.Port
		want map[string]*ports.Port
	}{
		"should create new port": {
			db: map[string]*ports.Port{},
			in: &ports.Port{
				Name:        "Ajman",
				City:        "Ajman",
				Country:     "United Arab Emirates",
				Coordinates: []float64{55.5136433, 25.4052165},
				Province:    "Ajman",
				Timezone:    "Asia/Dubai",
				Unlocs:      []string{"AEAJM"},
				Code:        "52000",
			},
			want: map[string]*ports.Port{
				"AEAJM": {
					Name:        "Ajman",
					City:        "Ajman",
					Country:     "United Arab Emirates",
					Coordinates: []float64{55.5136433, 25.4052165},
					Province:    "Ajman",
					Timezone:    "Asia/Dubai",
					Unlocs:      []string{"AEAJM"},
					Code:        "52000",
				},
			},
		},
		"should update port data": {
			db: map[string]*ports.Port{
				"AEAJM": {
					Name:        "Ajman",
					City:        "Ajman",
					Country:     "United Arab Emirates",
					Coordinates: []float64{55.5136433, 25.4052165},
					Province:    "Ajman",
					Timezone:    "Asia/Dubai",
					Unlocs:      []string{"AEAJM"},
					Code:        "52000",
				},
				"AEAUH": {
					Name:        "Abu Dhabi",
					City:        "Abu Dhabi",
					Country:     "United Arab Emirates",
					Coordinates: []float64{54.37, 24.47},
					Province:    "Ajman",
					Timezone:    "Asia/Dubai",
					Unlocs:      []string{"AEAUH"},
					Code:        "52001",
				},
			},
			in: &ports.Port{
				Name:        "Ajman",
				City:        "Ajman",
				Country:     "United Arab Emirates",
				Alias:       []string{"foo", "bar"},
				Coordinates: []float64{55.5136433, 25.4052165},
				Province:    "Ajman",
				Timezone:    "Asia/Dubai",
				Unlocs:      []string{"AEAJM"},
				Code:        "52000",
			},
			want: map[string]*ports.Port{
				"AEAJM": {
					Name:        "Ajman",
					City:        "Ajman",
					Country:     "United Arab Emirates",
					Alias:       []string{"foo", "bar"},
					Coordinates: []float64{55.5136433, 25.4052165},
					Province:    "Ajman",
					Timezone:    "Asia/Dubai",
					Unlocs:      []string{"AEAJM"},
					Code:        "52000",
				},
				"AEAUH": {
					Name:        "Abu Dhabi",
					City:        "Abu Dhabi",
					Country:     "United Arab Emirates",
					Coordinates: []float64{54.37, 24.47},
					Province:    "Ajman",
					Timezone:    "Asia/Dubai",
					Unlocs:      []string{"AEAUH"},
					Code:        "52001",
				},
			},
		},
	}
	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			store := memo{
				db: tc.db,
			}
			store.CreateOrUpdate(tc.in)
			if !reflect.DeepEqual(store.db, tc.want) {
				t.Fatalf("got %v; want %v", store.db, tc.want)
			}
		})
	}
}
