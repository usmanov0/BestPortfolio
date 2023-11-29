package error

import "fmt"

type RateLimitError struct {
	Ms int64
}

func (e *RateLimitError) Error() string {
	return fmt.Sprintf("Rate limit exceeded. Retry after %d milliseconds.", e.Ms)
}
