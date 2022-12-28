package uid

import "math/rand"

func CreateUid() int {
	return rand.Intn(1000000)
}
