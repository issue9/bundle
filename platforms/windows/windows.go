// SPDX-License-Identifier: MIT

// Package windows windows 平台实现
package windows

import (
	"errors"
	"go/build"
	"image/png"
	"os"
	"path/filepath"

	"github.com/issue9/bundle/internal/cmd"
	"github.com/issue9/bundle/platforms"
	xversion "github.com/issue9/version"
	"github.com/tc-hib/winres"
	"github.com/tc-hib/winres/version"
	"golang.org/x/text/language"
	"golang.org/x/text/message"
)

var archs = map[string]winres.Arch{
	"386":   winres.ArchI386,
	"amd64": winres.ArchAMD64,
	"arm":   winres.ArchARM,
	"arm64": winres.ArchARM64,
}

// targetDir 表示打包生成的应用保存的目录；
func Build(base *platforms.Options, o *Options) error {
	arch := build.Default.GOARCH
	src := filepath.Join(base.Root, base.Source)
	out := filepath.Join(base.Root, base.Output)

	syso := filepath.Join(src, arch+".syso")
	if err := writeResources(base, o, syso, arch); err != nil {
		return err
	}

	binPath := filepath.Join(out, base.Name+".exe")
	if err := cmd.Compile(src, binPath, "-ldflags", "-H=windowsgui"); err != nil {
		return err
	}

	return os.Remove(syso)
}

func writeResources(base *platforms.Options, o *Options, path, arch string) error {
	res := &winres.ResourceSet{}

	if err := buildIcon(base, res); err != nil {
		return err
	}

	if err := buildManifest(base, o, res); err != nil {
		return err
	}

	if err := buildVersion(base, o, res); err != nil {
		return err
	}

	out, err := os.Create(path)
	if err != nil {
		return err
	}
	defer out.Close()

	a, found := archs[arch]
	if !found {
		return errors.New("不支持的 arch:" + arch)
	}

	return res.WriteObject(out, a)
}

func buildVersion(base *platforms.Options, o *Options, res *winres.ResourceSet) error {
	v := version.Info{}

	for _, tag := range o.Languages {
		p := message.NewPrinter(tag, message.Catalog(o.Catalog))

		var lcid uint16 = version.LangNeutral
		if tag != language.Und {
			id, err := getWindowsLCID(tag)
			if err != nil {
				return err
			}

			lcid = uint16(id)
		}

		v.Set(lcid, version.LegalCopyright, o.Copyright.LocaleString(p))
		v.Set(lcid, version.Comments, o.Comments.LocaleString(p))
		v.Set(lcid, version.FileDescription, o.FileDescription.LocaleString(p))
		v.Set(lcid, version.ProductName, o.ProductName.LocaleString(p))
		if o.CompanyName != nil {
			v.Set(lcid, version.CompanyName, o.CompanyName.LocaleString(p))
		}
	}

	v.SetFileVersion(base.Version)
	v.SetProductVersion(base.Version)

	res.SetVersionInfo(v)
	return nil
}

func buildManifest(base *platforms.Options, o *Options, res *winres.ResourceSet) error {
	v := &xversion.SemVersion{}
	if err := xversion.Parse(v, base.Version); err != nil {
		return err
	}

	res.SetManifest(winres.AppManifest{
		Identity: winres.AssemblyIdentity{
			Name:    base.ID,
			Version: [4]uint16{uint16(v.Major), uint16(v.Minor), uint16(v.Patch), 0},
		},
		Compatibility:       o.MinOS,
		ExecutionLevel:      winres.AsInvoker,
		DPIAwareness:        winres.DPIPerMonitorV2,
		UseCommonControlsV6: true,
	})
	return nil
}

func buildIcon(base *platforms.Options, res *winres.ResourceSet) error {
	r, err := os.Open(filepath.Join(base.Root, base.Icon))
	if err != nil {
		return err
	}
	defer r.Close()

	img, err := png.Decode(r)
	if err != nil {
		return err
	}
	icon, err := winres.NewIconFromResizedImage(img, []int{16, 24, 32, 64, 128, 256})
	if err != nil {
		return err
	}

	return res.SetIcon(winres.Name("APPICON"), icon)
}
