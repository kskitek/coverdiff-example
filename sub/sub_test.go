package sub_test

import (
	"fmt"
	"testing"

	"github.com/kskitek/coverdiff-example/sub"
)

func TestSub(t *testing.T) {
	cases := []struct {
		a        int
		b        int
		expected int
	}{
		{3, 2, 1},
	}

	for i, tc := range cases {
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			result := sub.Sub(tc.a, tc.b)
			if result != tc.expected {
				t.Errorf("expected %d got %d", tc.expected, result)
			}
		})
	}
}
