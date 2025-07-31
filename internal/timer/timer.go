package timer

import (
	"fmt"
	"time"
)

func Timer(duration time.Duration) {
	seconds := int(duration.Seconds())
	regressingDurat := seconds

	for i := 0; i <= seconds; i++ {
		actualTime := time.Duration(regressingDurat) * time.Second
		fmt.Printf("\r%s restantes ", actualTime)
		time.Sleep(1 * time.Second)
		regressingDurat--
	}
}
