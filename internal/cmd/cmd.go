// SPDX-License-Identifier: MIT

package cmd

import (
	"errors"
	"fmt"
	"io/fs"
	"os"
	"os/exec"
	"path/filepath"
)

func Command(name string, arg ...string) *exec.Cmd {
	c := exec.Command(name, arg...)
	c.Stderr = os.Stderr
	c.Stdout = os.Stdout
	return c
}

// 执行一组命令行
func Commands(cmds ...*exec.Cmd) error {
	for _, c := range cmds {
		if err := c.Run(); err != nil {
			return err
		}
	}
	return nil
}

// Copy 复制文件
//
// 如果 dest 为目录，则会将 src 文件作为其下的子文件；
// 如果 src 为目录，则返回错误；
func Copy(src, dest string) error {
	stat, err := os.Stat(src)
	if err != nil {
		return err
	}
	if stat.IsDir() {
		return &fs.PathError{
			Op:   "copyy",
			Path: "src",
			Err:  errors.New("只能是文件"),
		}
	}

	data, err := os.ReadFile(src)
	if err != nil {
		return err
	}

	stat, err = os.Stat(dest)
	if err != nil {
		if errors.Is(err, fs.ErrNotExist) {
			return os.WriteFile(dest, data, os.ModePerm)
		}
		return err
	} else if stat.IsDir() { // 目录
		return os.WriteFile(filepath.Join(dest, filepath.Base(src)), data, os.ModePerm)
	} else { // 非目录
		return os.WriteFile(dest, data, os.ModePerm)
	}
}

func Move(src, dest string) error {
	if err := Copy(src, dest); err != nil {
		return err
	}
	return os.Remove(src)
}

// Compile 编译 go 代码
//
// main 为源码目录；
// out 为输出文件；
// arg 为其它参数；
func Compile(main, out string, arg ...string) (err error) {
	fmt.Println("开始编译二进制文件...")
	defer func() {
		if err == nil {
			fmt.Println("编译完成...")
		}
	}()

	args := []string{
		"build",
		"-v",
		"-o", out,
	}
	args = append(args, arg...)
	args = append(args, main)
	err = Command("go", args...).Run() // defer 需要 err 变量
	return err
}
