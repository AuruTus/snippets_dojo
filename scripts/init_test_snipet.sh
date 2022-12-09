#!/bin/bash

TSTR_FILE=$1
TSTR=$2

cat > ./src/test_snipets/${TSTR_FILE} << EOF
package test_snipets

import (
	"context"
	snipets "snipets_dojo/src"
)

type $TSTR struct{}

var _ snipets.Tstr = (*$TSTR)(nil)

func (tstr *$TSTR) Test(ctx context.Context) error {
	return nil
}


EOF