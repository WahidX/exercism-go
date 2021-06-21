package gigasecond

import (
	"math"
	"time"
)

func AddGigasecond(t time.Time) (ret time.Time) {
	var duration time.Duration = time.Second*time.Duration(math.Pow(10,9))

	ret = t.Add(duration)

	return 
}
