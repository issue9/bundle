// SPDX-License-Identifier: MIT

// Package bundle Go 语言的打包工具
package bundle

import (
	"go/build"

	"github.com/issue9/bundle/platforms"
	"github.com/issue9/bundle/platforms/macos"
	"github.com/issue9/bundle/platforms/windows"
)

type Options struct {
	Base *platforms.Options

	MacOS *macos.Options

	Windows *windows.Options
}

// Build 编译为指定平台的 GUI 程序
func Build(o *Options) error {
	switch build.Default.GOOS {
	case "windows":
		return windows.Build(o.Base, o.Windows)
	case "darwin":
		return macos.Build(o.Base, o.MacOS)
	default:
		return platforms.Build(o.Base)
	}
}
