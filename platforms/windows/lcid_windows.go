// SPDX-License-Identifier: MIT

package windows

import (
	"github.com/issue9/localeutil/windows"
	"golang.org/x/text/language"
)

func getWindowsLCID(tag language.Tag) (uint32, error) { return windows.GetLCID(tag.String()) }
