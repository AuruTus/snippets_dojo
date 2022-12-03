package test_snipets

import "context"

type Tstr interface {
	Test(context.Context) error
}
