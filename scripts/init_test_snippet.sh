#!/bin/bash

TSTR_FILE=$1
TSTR=$2

cat > ./src/test_snippets/${TSTR_FILE} << EOF
package test_snippets

import (
	"context"
	snippets "snippets_dojo/src"
)

type $TSTR struct{}

var _ snippets.Tstr = (*$TSTR)(nil)

func (tstr *$TSTR) Test(ctx context.Context) error {
	return nil
}


EOF