package snipets

import "context"

type Tstr interface {
	Test(context.Context) error
}
