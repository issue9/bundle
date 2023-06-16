// SPDX-License-Identifier: MIT

package platforms

// Options 基本的打包选项
type Options struct {
	// Root 项目根目录
	Root string

	// Source 需要编译的源码目录
	//
	// 相对于 Root 的路径
	Source string

	// Output 编译后的输出目录
	//
	// 相对于 Root 的路径
	Output string

	// Icon 应用图标
	//
	// 一个尺寸足够大的 png 格式图片，相对于 Root 的路径。
	Icon string

	// Name 应用名称
	Name string

	// Version 应用版本
	Version string

	// ID 应用的唯一 ID
	ID string
}
