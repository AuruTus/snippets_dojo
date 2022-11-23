package snipets

import (
	"context"
)

type RWQLockTstr struct{}

func (tstr *RWQLockTstr) Test(ctx context.Context) error {
	return nil
}
