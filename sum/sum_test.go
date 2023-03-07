package sum_test

import (
	"fmt"
	"testing"

	"github.com/kskitek/coverdiff-example/sum"
)

func TestSum(t *testing.T) {
	cases := []struct {
		a        int
		b        int
		expected int
	}{
		// {1, 2, 3},
	}

	for i, tc := range cases {
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			result := sum.Sum(tc.a, tc.b)
			if result != tc.expected {
				t.Errorf("expected %d got %d", tc.expected, result)
			}
		})
	}
}
