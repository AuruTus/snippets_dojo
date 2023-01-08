#!/bin/bash

TSTR_FILE=$1
TSTR=$2

# Create new tester file and tstr struct
cat > ./src/test_snippets/${TSTR_FILE} << EOF
package test_snippets

import (
	"context"
	snippets "snippets_dojo/src"
)

type $TSTR struct{}

var _ snippets.Tstr = (*$TSTR)(nil)

func (tstr *$TSTR) Test(ctx context.Context) error {
	log := cfmt.Logger(ctx)
	return nil
}

EOF

# Change the test entry
cat > ./tstr_entry.go << EOF
package main

import (
	snippets "snippets_dojo/src"
	test_snippets "snippets_dojo/src/test_snippets"
)

func NewTesterEntry() snippets.Tstr {
	// !NOTE: customize tester entrance here
	return &test_snippets.$TSTR{}
}

EOF