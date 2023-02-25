package sum

import "testing"

/* name 			input	        want
one to five 	[]int{1...5}     15
nil slice       nil 			 0

*/

func TestSumInt(t *testing.T) {
	tt := []struct {
		//fields
		name    string
		numbers []int
		want    int
	}{
		{
			name:    "one to five",
			numbers: []int{1, 2, 3, 4, 5},
			want:    15,
		},
		{
			name:    "nil slice",
			numbers: nil,
			want:    0,
		},
		{
			name:    "one minus one",
			numbers: []int{1, -1},
			want:    0,
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) { // it creates a subtest // we can run each struct instance as an individual test case
			got := SumInt(tc.numbers)
			if got != tc.want {
				t.Errorf("sum of %v want %v; got %v", tc.numbers, tc.want, got)
			}
		})
	}
}
