package ctxinfo

import (
	"context"
	"fmt"

	"github.com/google/uuid"
)

type key int

const (
	CTX_INFO_KEY key = iota
)

type ContextLevel int32

const (
	BACKGROUND ContextLevel = iota
	MAIN
	TESTER
)

type ContextInfo struct {
	CtxID uuid.UUID
	Level ContextLevel
}

func NewContextWithInfo(
	bg_ctx context.Context,
	ctxLevel ContextLevel,
) (ctx context.Context,
	cancel context.CancelFunc) {
	ctx = context.WithValue(
		bg_ctx,
		CTX_INFO_KEY,
		ContextInfo{
			CtxID: uuid.New(),
			Level: ctxLevel,
		},
	)
	ctx, cancel = context.WithCancel(ctx)
	return
}

func Unwrap(info any) (ContextInfo, error) {
	switch data := info.(type) {
	case ContextInfo:
		return data, nil
	case *ContextInfo:
		if data == nil {
			return ContextInfo{}, fmt.Errorf("empty context info")
		}
		return *data, nil
	default:
		return ContextInfo{}, fmt.Errorf("invalid context info")
	}
}
