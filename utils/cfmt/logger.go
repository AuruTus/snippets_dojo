package cfmt

import (
	"context"
	"fmt"
	"log"
	"os"
)

// TODO
// 1. a logger init entrance
// 2. customizable format including file line
//

func InitLogger() {
	log.SetOutput(os.Stdout)
	log.SetFlags(log.Lshortfile)
}

type LogWrapper struct {
	Log func(format string, args ...any) (int, error)
}

func NewLogger(ctx context.Context) *LogWrapper {
	return &LogWrapper{
		Log: func(format string, args ...any) (int, error) {
			return Printf(ctx, fmt.Sprintf("%s\n", format), args)
		},
	}
}
