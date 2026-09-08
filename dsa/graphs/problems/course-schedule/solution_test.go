package courseschedule

import "testing"

func TestCanFinish(t *testing.T) {
	tests := []struct {
		name          string
		numCourses    int
		prerequisites [][2]int
		want          bool
	}{
		{"no prerequisites", 2, [][2]int{}, true},
		{"simple chain", 2, [][2]int{{1, 0}}, true},
		{"direct cycle", 2, [][2]int{{1, 0}, {0, 1}}, false},
		{"self loop", 1, [][2]int{{0, 0}}, false},
		{"diamond dependency", 4, [][2]int{{1, 0}, {2, 0}, {3, 1}, {3, 2}}, true},
		{"cycle in larger graph", 4, [][2]int{{1, 0}, {2, 1}, {0, 2}, {3, 2}}, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CanFinish(tt.numCourses, tt.prerequisites); got != tt.want {
				t.Errorf("CanFinish(%d, %v) = %v, want %v",
					tt.numCourses, tt.prerequisites, got, tt.want)
			}
		})
	}
}
