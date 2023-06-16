// SPDX-License-Identifier: MIT

package windows

import (
	"github.com/issue9/localeutil"
	"github.com/tc-hib/winres"
	"golang.org/x/text/language"
	"golang.org/x/text/message/catalog"
)

type Options struct {
	// Languages 支持的语言列表
	//
	// 如果值中包含 [language.Und]，将会被映射为 windows 的 Neutral。
	Languages []language.Tag

	// MinOS 支持最低操作系统
	MinOS SupportedOS

	Catalog catalog.Catalog

	Copyright       localeutil.LocaleStringer
	Comments        localeutil.LocaleStringer
	FileDescription localeutil.LocaleStringer
	ProductName     localeutil.LocaleStringer
	CompanyName     localeutil.LocaleStringer // 可以为 nil
}

type SupportedOS = winres.SupportedOS

const (
	Win7AndAbove  = winres.Win7AndAbove
	Win8AndAbove  = winres.Win8AndAbove
	Win81AndAbove = winres.Win81AndAbove
	Win10AndAbove = winres.Win10AndAbove
)
