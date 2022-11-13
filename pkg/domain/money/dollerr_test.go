package money

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestTimes(t *testing.T) {
	cases := map[string]struct {
		doller  *Doller
		multiplier  int
		want    *Doller
	}{
		"5ドルに2を乗算して10が返却される": {
			doller:   NewDoller(5),
			multiplier: 2,
			want: NewDoller(10),
		},
		"2ドルに6を乗算して12が返却される": {
			doller:   NewDoller(2),
			multiplier: 6,
			want: NewDoller(12),
		},
	}

	for name, tc := range cases {
		t.Run(name, func(t *testing.T) {
			res := tc.doller.times(tc.multiplier)

			if diff := cmp.Diff(res, tc.want); diff != "" {
				t.Errorf("X value is mismatch (-want +got):%s\n", diff)
			}

		})
	}
}
