// SPDX-License-Identifier: MIT

// Package macos 用于打包 macOS 应用
package macos

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/issue9/errwrap"
	"github.com/issue9/localeutil"
	"golang.org/x/text/message"
	"howett.net/plist"

	"github.com/issue9/bundle/internal/cmd"
	"github.com/issue9/bundle/internal/preset"
	"github.com/issue9/bundle/platforms"
)

// https://developer.apple.com/documentation/bundleresources/information_property_list/bundle_configuration
type infoPlist struct {
	Name         string   `plist:"CFBundleName"`
	DisplayName  string   `plist:"CFBundleDisplayName,omitempty"`
	Icon         string   `plist:"CFBundleIconFile"`
	ID           string   `plist:"CFBundleIdentifier"`
	Version      string   `plist:"CFBundleVersion"`
	ShortVersion string   `plist:"CFBundleShortVersionString,omitempty"`
	Copyright    string   `plist:"NSHumanReadableCopyright"`
	MinOS        string   `plist:"LSMinimumSystemVersion,omitempty"`
	Executable   string   `plist:"CFBundleExecutable"`
	Languages    []string `plist:"CFBundleLocalizations,omitempty"`
	PackageType  string   `plist:"CFBundlePackageType"`
	Category     string   `plist:"LSApplicationCategoryType"`
}

// Build 打包生成 macOS 的应用
func Build(base *platforms.Options, o *Options) error {
	output := filepath.Join(base.Root, base.Output)
	appDir := filepath.Join(output, base.Name+".app")
	contentDir := "Contents"
	resDir := filepath.Join(appDir, contentDir, "Resources")
	execDir := filepath.Join(appDir, contentDir, "MacOS")

	if err := os.MkdirAll(resDir, os.ModePerm); err != nil {
		return err
	}
	if err := os.MkdirAll(execDir, os.ModePerm); err != nil {
		return err
	}

	info := &infoPlist{
		Name:         base.Name,
		DisplayName:  base.Name,
		Icon:         base.Name,
		ID:           base.ID,
		Version:      base.Version,
		ShortVersion: base.Version,
		Copyright:    o.Copyright,
		MinOS:        o.MinOS,
		Executable:   base.Name,
		PackageType:  "APPL",
		Category:     string(o.Category),
	}
	fmt.Println("生成 info.plist")
	if err := info.makeLocales(o, resDir); err != nil {
		return err
	}
	data, err := plist.MarshalIndent(info, plist.XMLFormat, "    ")
	if err != nil {
		return err
	}
	if err := os.WriteFile(filepath.Join(appDir, contentDir, "Info.plist"), data, os.ModePerm); err != nil {
		return err
	}
	fmt.Println("成功生成 info.plist")

	iconPath := filepath.Join(resDir, base.Name+".icns")
	if err := makeIcns(base, iconPath, filepath.Join(base.Root, base.Icon)); err != nil {
		return err
	}

	binName := base.Name
	if err = cmd.Compile(filepath.Join(base.Root, base.Source), binName); err != nil {
		return err
	}
	return cmd.Move(binName, filepath.Join(execDir, base.Name))
}

// 生成本地化相关内容
//
// 包含了以下内容：
//   - info.plist 中的 CFBundleLocalizations 字段；
//   - CFBundleLocalizations 对应的本地化目录和文件；
func (info *infoPlist) makeLocales(o *Options, resDir string) error {
	for _, tag := range o.Languages {
		id := tag.String()
		info.Languages = append(info.Languages, id)

		dir := filepath.Join(resDir, id+".lproj")
		if err := os.MkdirAll(dir, os.ModePerm); err != nil {
			return err
		}

		data, err := buildStrings(o, message.NewPrinter(tag, message.Catalog(o.Catalog)))
		if err != nil {
			return err
		}

		if err := os.WriteFile(filepath.Join(dir, "InfoPlist.strings"), data, os.ModePerm); err != nil {
			return err
		}
	}

	return nil
}

func buildStrings(o *Options, p *localeutil.Printer) ([]byte, error) {
	w := errwrap.Buffer{}

	w.WString(preset.FileHeader("// "))

	w.Printf(`"NSHumanReadableCopyright"="%s";`, o.ReadableCopyright.LocaleString(p)).WString("\n\n")
	w.Printf(`"CFBundleDisplayName"="%s";`, o.DisplayName.LocaleString(p)).WString("\n\n")
	w.Printf(`"CFBundleGetInfoString"="%s";`, o.GetInfoString.LocaleString(p)).WString("\n\n")

	return w.Bytes(), w.Err
}

// 生成 icns 的图标
//
// target 为生成的图标保存的地址，包含了文件名；
// png 为图标的源文件，将根据此文件生成 icon.icns 图标文件；
func makeIcns(base *platforms.Options, target, png string) error {
	fmt.Println("开始生成 icns")
	defer fmt.Println("成功生成 icns 文件")

	dir, err := os.MkdirTemp("", base.Name)
	if err != nil {
		return err
	}
	set := filepath.Join(dir, "macos.iconset")
	if err := os.MkdirAll(set, os.ModePerm); err != nil {
		return err
	}
	out := filepath.Join(dir, "icon.icns")

	err = cmd.Commands(
		cmd.Command("sips", "-z", "16", "16", png, "-o", filepath.Join(set, "icon_16x16.png")),
		cmd.Command("sips", "-z", "32", "32", png, "-o", filepath.Join(set, "icon_16x16@2x.png")),
		cmd.Command("sips", "-z", "32", "32", png, "-o", filepath.Join(set, "icon_32x32.png")),
		cmd.Command("sips", "-z", "64", "64", png, "-o", filepath.Join(set, "icon_32x32@2x.png")),
		cmd.Command("sips", "-z", "64", "64", png, "-o", filepath.Join(set, "icon_64x64.png")),
		cmd.Command("sips", "-z", "128", "128", png, "-o", filepath.Join(set, "icon_64x64@2x.png")),
		cmd.Command("sips", "-z", "128", "128", png, "-o", filepath.Join(set, "icon_128x128.png")),
		cmd.Command("sips", "-z", "256", "256", png, "-o", filepath.Join(set, "icon_128x128@2x.png")),
		cmd.Command("sips", "-z", "256", "256", png, "-o", filepath.Join(set, "icon_256x256.png")),
		cmd.Command("sips", "-z", "512", "512", png, "-o", filepath.Join(set, "icon_256x256@2x.png")),
		cmd.Command("sips", "-z", "512", "512", png, "-o", filepath.Join(set, "icon_512x512.png")),
		cmd.Command("sips", "-z", "1024", "1024", png, "-o", filepath.Join(set, "icon_512x512@2x.png")),
		cmd.Command("iconutil", "-c", "icns", set, "-o", out),
	)
	if err != nil {
		return err
	}

	return cmd.Move(out, target)
}
