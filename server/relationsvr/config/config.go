package config

import (
	"os"
	"path/filepath"
)

var rootStr string

// GetRootStr 获取项目主目录
func GetRootStr() string {
	return rootStr
}

func init() {
	setRootStr()
}

func setRootStr() {
	dir, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	var infer func(dir string) string
	infer = func(dir string) string {
		if exist(dir + "/main.go") {
			return dir
		}
		// 查看dir的父目录
		parent := filepath.Dir(dir)
		return infer(parent)
	}
	rootStr = infer(dir)
}

func exist(dir string) bool {
	_, err := os.Stat(dir)
	return err == nil || os.IsExist(err)
}
