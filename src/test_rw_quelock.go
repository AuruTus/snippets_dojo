package snipets

import (
	"context"
)

type RWQLockTstr struct{}

var _ Tstr = (*RWQLockTstr)(nil)

func (tstr *RWQLockTstr) Test(ctx context.Context) error {
	return nil
}
