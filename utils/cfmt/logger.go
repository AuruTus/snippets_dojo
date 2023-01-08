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

func Logger(ctx context.Context) func(format string, args ...any) (int, error) {
	simpleLogger := func(format string, args ...any) (int, error) {
		return Printf(ctx, fmt.Sprintf("%s\n", format), args)
	}
	return simpleLogger
}
