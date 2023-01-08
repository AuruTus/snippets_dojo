package cfmt

import (
	"context"
	"fmt"
)

type LogWrapper struct {
	Log func(format string, args ...any) (int, error)
}

func NewSimpleLogger(ctx context.Context) *LogWrapper {
	return &LogWrapper{
		Log: func(format string, args ...any) (int, error) {
			return Printf(ctx, fmt.Sprintf("%s\n", format), args...)
		},
	}
}
