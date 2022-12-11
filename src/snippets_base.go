package test_snippets

import (
	"context"
	"snippets_dojo/utils/cfmt"
)

type Tstr interface {
	Test(context.Context) error
}

/**
 * Dummy tester for tstr_entry.go code generator place holder.
 */
type DummyTstr struct{}

var _ Tstr = (*DummyTstr)(nil)

func (t *DummyTstr) Test(ctx context.Context) error {
	cfmt.Printf(ctx, "%s %s\n", cfmt.Teal("Hello"), cfmt.Purple("Dojo"))
	return nil
}
