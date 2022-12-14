package cfmt

import (
	"context"
	"fmt"
	"runtime"
	ctxinfo "snippets_dojo/utils/ctx_info"
	"strings"
)

/*
* ** package private constants
 */
const (
	_CALLER_STACK_SKIP_DEPTH = 3

	_MODULE_NAME = "snippets_dojo"
)

func getCallerLine(skip int) (file string, line int) {
	_, file, line, ok := runtime.Caller(skip)
	if !ok {
		file = "???"
		line = 0
		return
	}

	// if the project's root directory name is changed, don't print
	// relative path.
	index := strings.Index(file, _MODULE_NAME)
	if index != -1 {
		file = "." + file[index+len(_MODULE_NAME):]
	}
	return
}

func baseCallerLine() (file string, line int) {
	return getCallerLine(_CALLER_STACK_SKIP_DEPTH)
}

func simplePrintf(ctx context.Context, format string, args ...interface{}) (int, error) {
	file, line := baseCallerLine()
	return fmt.Printf(
		"%s: %s",
		Green(fmt.Sprintf("%s %d", file, line)),
		fmt.Sprintf(format, args...),
	)
}

func verbosePrintf(ctx context.Context, format string, args ...interface{}) (int, error) {
	info, _ := ctxinfo.Unwrap(ctx.Value(ctxinfo.CTX_INFO_KEY))

	file, line := baseCallerLine()

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
