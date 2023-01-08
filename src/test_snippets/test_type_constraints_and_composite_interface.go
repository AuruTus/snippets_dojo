package test_snippets

import (
	"context"
	snippets "snippets_dojo/src"
	"snippets_dojo/utils/cfmt"
)

type IntrTstr struct{}

var _ snippets.Tstr = (*IntrTstr)(nil)

type Cow interface {
	Mow() int64
}

type Duck interface {
	Quack() bool
}

/*
 * ===================
 * Composite Interface
 * ===================
 */
type CowDuck interface {
	Cow
	Duck
}

type CDStruct struct{}

func (cd CDStruct) Mow() int64  { return 1 }
func (cd CDStruct) Quack() bool { return false }

var _ CowDuck = (*CDStruct)(nil)
var _ Cow = (*CDStruct)(nil)
var _ Duck = (*CDStruct)(nil)

/*
 * ===================
 * Type Constraints
 * ===================
 */
type CowMow int64
type DuckQuack bool

type CowDuckConstraints interface {
	CowMow | DuckQuack
}

/*
 * ===================
 * Type Constraints mixes Composite Interface
 * ===================
 */
type CowDuckMixed interface {
	CDStruct | CowDuckConstraints

	CowDuck
}

func (i CowMow) Mow() int64  { return 1 }
func (i CowMow) Quack() bool { return false }

/* invalid */
// var _ CowDuckMixed = (*CDStruct)(nil)

func testMixed[T CowDuckMixed](ctx context.Context, v T) {
	cfmt.Printf(ctx, "mixed type constraints and interface: %T\n", v)
}

func (tstr *IntrTstr) Test(ctx context.Context) error {
	testMixed(ctx, CDStruct{})
	testMixed[CowMow](ctx, 1)
	return nil
}
