package knownpattern

import (
	"fmt"
	"time"
)

func Retry(operation func() error, retries int) error {
	for i := 0; i < retries; i++ {
		if err := operation(); err != nil {
			time.Sleep(time.Second << i)
			continue
		}
		return nil
	}
	return fmt.Errorf("operation failed after %d retries", retries)
}
