package dailytemperatures

import (
	"reflect"
	"testing"
)

func TestDailyTemperatures(t *testing.T) {
	tests := []struct {
		name         string
		temperatures []int
		want         []int
	}{
		{"classic example", []int{73, 74, 75, 71, 69, 72, 76, 73}, []int{1, 1, 4, 2, 1, 1, 0, 0}},
		{"strictly decreasing", []int{5, 4, 3, 2, 1}, []int{0, 0, 0, 0, 0}},
		{"strictly increasing", []int{1, 2, 3, 4}, []int{1, 1, 1, 0}},
		{"single day", []int{50}, []int{0}},
		{"all equal", []int{60, 60, 60}, []int{0, 0, 0}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := DailyTemperatures(tt.temperatures); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("DailyTemperatures(%v) = %v, want %v", tt.temperatures, got, tt.want)
			}
		})
	}
}
