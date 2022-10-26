//-----------------------------------------------------------------------------
// Copyright (c) 2022 Detlef Stern
//
// This file is part of gopikchr.
//-----------------------------------------------------------------------------

package gopikchr_test

import (
	"testing"

	"github.com/t73fde/gopikchr"
)

func FuzzConvert(f *testing.F) {
	f.Fuzz(func(t *testing.T, src []byte) {
		t.Parallel()
		gopikchr.Convert(src)
	})
}
