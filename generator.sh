#!/bin/bash

TSTR_ENTRY=./tstr_entry.go

cat > ${TSTR_ENTRY} << EOF
package main

import (
	snipets "snipets_dojo/src"
)

func NewTesterEntry() snipets.Tstr {
	// !NOTE: customize tester entrance here
	return &snipets.DummyTstr{}
}

EOF