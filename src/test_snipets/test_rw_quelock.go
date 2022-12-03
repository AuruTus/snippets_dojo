package test_snipets

import (
	"context"
	snipets "snipets_dojo/src"
)

type RWQLockTstr struct{}

var _ snipets.Tstr = (*RWQLockTstr)(nil)

func (tstr *RWQLockTstr) Test(ctx context.Context) error {
	return nil
}
