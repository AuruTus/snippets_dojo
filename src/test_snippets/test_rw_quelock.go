package test_snippets

import (
	"context"
	snippets "snippets_dojo/src"
)

type RWQLockTstr struct{}

var _ snippets.Tstr = (*RWQLockTstr)(nil)

func (tstr *RWQLockTstr) Test(ctx context.Context) error {
	return nil
}
