package helpers

import (
	"time"
)

func IntervalFunction(callback func(), interval time.Duration) {
	ticker := time.NewTicker(interval * time.Second)
	quit := make(chan struct{})

	go func() {
		for {
			select {
			case <-ticker.C:
				callback()
			case <-quit:
				ticker.Stop()
				return
			}
		}
	}()
}
