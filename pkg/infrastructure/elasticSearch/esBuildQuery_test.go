package elasticsearch

import (
	"testing"
)

func TestNewbBildQuery(t *testing.T) {
	t.Run("正常", func(t *testing.T) {
		buildQuery([]string{"test"}, 10, true)
	})
}
