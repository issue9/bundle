// SPDX-License-Identifier: MIT

// Package platforms 平台相关在的打包方法
package platforms

import (
	"path/filepath"

	"github.com/issue9/bundle/internal/cmd"
)

func Build(o *Options) error {
	output := filepath.Join(o.Root, o.Output, o.Name)
	src := filepath.Join(o.Root, o.Source)
	return cmd.Compile(src, output)
}
