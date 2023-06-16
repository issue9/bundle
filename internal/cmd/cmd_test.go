// SPDX-License-Identifier: MIT

package cmd

import (
	"os"
	"testing"

	"github.com/issue9/assert/v3"
)

func TestCopy_Move(t *testing.T) {
	a := assert.New(t, false)

	a.Error(Copy("./", "./testdata"))
	a.NotError(Copy("./cmd.go", "./testdata"))
	a.NotError(Copy("./cmd.go", "./testdata/cmd.go")) // 覆盖
	a.NotError(Copy("./cmd.go", "./testdata/icon1.png"))

	a.NotError(Move("testdata/cmd.go", "./testdata/icon2.png"))

	a.NotError(os.Remove("./testdata/icon2.png"))
	a.NotError(os.Remove("./testdata/icon1.png"))
}
