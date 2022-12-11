#!/bin/bash

TSTR_ENTRY=./tstr_entry.go

cat > ${TSTR_ENTRY} << EOF
package main

import (
	snippets "snippets_dojo/src"
)

func NewTesterEntry() snippets.Tstr {
	// !NOTE: customize tester entrance here
	return &snippets.DummyTstr{}
}

EOF