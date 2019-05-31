package maximum

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestFindMaximum(t *testing.T) {
	r := require.New(t)

	var tt = map[string]struct {
		in    []float64
		max1  float64
		max2  float64
		error error
	}{
		"same numbers": {
			[]float64{1, 1.0, 1, 1}, 0, 0, ErrNoMaximum,
		},
		"basic example": {
			[]float64{4, 23, 5, 23, 7}, 7, 23 , nil,
		},
		"first elements are same": {
			[]float64{4, 4, 1, 3}, 3,4,nil,
		},
		"last elements are same": {
			[]float64{1, 3, 4,4}, 3,4,nil,
		},
	}

	for name, tc := range tt {
		t.Run(name, func(t *testing.T) {
			max1, max2, err := FindMaximum(tc.in)
			r.Equal(tc.max1, max1)
			r.Equal(tc.max2, max2)
			r.EqualValues(tc.error, err)
		})
	}
}
