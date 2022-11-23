package cfmt

import (
	"context"
	"fmt"
	"runtime"
	ctxinfo "snipets_dojo/utils/ctx_info"
	"strings"
)

/*
* ** package private constants
 */
const (
	_CALLER_STACK_SKIP_DEPTH = 3

	_MODULE_NAME = "test_snipets"
)

func getCallerLine(skip int) (file string, line int) {
	_, file, line, ok := runtime.Caller(_CALLER_STACK_SKIP_DEPTH)
	if !ok {
		file = "???"
		line = 0
	} else {
		index := strings.Index(file, _MODULE_NAME)
		file = "." + file[index+len(_MODULE_NAME):]
	}
	return file, line
}

func simplePrintf(ctx context.Context, format string, args ...interface{}) (int, error) {
	file, line := getCallerLine(_CALLER_STACK_SKIP_DEPTH)
	return fmt.Printf(
		"%s: %s",
		Green(fmt.Sprintf("%s %d", file, line)),
		fmt.Sprintf(format, args...),
	)
}

func verbosePrintf(ctx context.Context, format string, args ...interface{}) (int, error) {
	info, _ := ctxinfo.Unwrap(ctx.Value(ctxinfo.CTX_INFO_KEY))

	file, line := getCallerLine(_CALLER_STACK_SKIP_DEPTH)

	return fmt.Printf(
		"@%s | %s | %s",
		Green(fmt.Sprintf("%s %d", file, line)),
		Green(info.CtxID.String()),
		fmt.Sprintf(format, args...),
	)
}

func Printf(ctx context.Context, format string, args ...interface{}) (int, error) {
	info, _ := ctxinfo.Unwrap(ctx.Value(ctxinfo.CTX_INFO_KEY))
	switch info.Level {
	case ctxinfo.TESTER:
		return simplePrintf(ctx, format, args...)
	default:
		// for background and main context
		return verbosePrintf(ctx, format, args...)
	}
}
