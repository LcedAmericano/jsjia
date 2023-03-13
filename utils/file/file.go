/*
Package file 与文件相关的工具包
*/
package file

import "os"

// Exists 判断文件或文件夹是否存在
func Exists(name string) bool {
	if _, err := os.Stat(name); err != nil {
		if os.IsNotExist(err) {
			return false
		}
	}
	return true
}
