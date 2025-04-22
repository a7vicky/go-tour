package knownpattern

import "errors"

var failureCount int

func CircuitBreaker(operation func() error, threshold int) error {
	if failureCount >= threshold {
		return errors.New("circuit breaker open")
	}

	err := operation()
	if err != nil {
		failureCount++
		return err
	}

	failureCount = 0
	return nil
}
