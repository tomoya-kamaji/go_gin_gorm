package errorCheck

import (
	"testing"
	"yu-croco/ddd_on_golang/pkg/errors"

	"github.com/google/go-cmp/cmp"
)

func Check(t testing.TB, err *errors.AppError, wantErr *errors.AppError) bool {
	t.Helper()

	// エラーが発生していないが、期待しているエラーがある場合
	if !err.HasErrors() {
		if wantErr.HasErrors() {
			t.Errorf("error was expceted but no error was returned:\n    wantErr: %v", wantErr)
		}
		return true
	}

	// エラーが発生していて、内容に相違がある場合
	if wantErr.HasErrors() {
		if diff := cmp.Diff(err, wantErr); diff != "" {
			t.Errorf("error did not match:\n    want: %v\n    got:  %v", wantErr, err)
		}
		return false
	}

	t.Errorf("error was not expected got one:\n    got: %v", err)
	return false
}
