// SPDX-License-Identifier: MIT

// Package preset 一些默认设置
package preset

const fileHeader = "此文件由 github.com/issue9/bundle 生成，请勿手动修改！"

// FileHeader 生成由工具生成的代码文件头
//
//	prefix 表示注释字符，比如 go 中为 //，sh 中为 #
func FileHeader(prefix string) string { return prefix + fileHeader + "\n\n" }
