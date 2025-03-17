package bin

import "time"

type Bins []*Bin

func (b Bins) Next() Bin {
	var closestFutureBin Bin
	var minDuration time.Duration

	now := time.Now()

	for _, bin := range b {
		if bin.Next.After(now) {
			nextbinDate := bin.Next

			duration := nextbinDate.Sub(now)
			if closestFutureBin.Next.IsZero() || duration < minDuration {
				closestFutureBin = *bin
				minDuration = duration
			}
		}
	}

	return closestFutureBin
}
