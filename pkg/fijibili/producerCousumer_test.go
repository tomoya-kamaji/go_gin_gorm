package parallelprocessing

import (
	"testing"
)

func TestTimes(t *testing.T) {
	t.Run("テスト", func(t *testing.T) {
		print(t)
		producerCousumer()
	})
}
