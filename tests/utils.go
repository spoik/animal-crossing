package tests

import (
	"math/rand"
	"time"
)

func RandomTime() time.Time {
	randomTime := rand.Int63n(time.Now().Unix() - 94608000) + 94608000
	return time.Unix(randomTime, 0)
}
