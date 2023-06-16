// SPDX-License-Identifier: MIT

package windows

import (
	"fmt"
	"runtime"

	"golang.org/x/text/language"
)

// 该函数仅为了编译不出错，并无实现用处。
func getWindowsLCID(tag language.Tag) (uint32, error) {
	return 0, fmt.Errorf("当前系统 %s 不需要", runtime.GOOS)
}
