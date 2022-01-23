package mian

import "os"

func PathIfExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		// 文件或目录存在
		return true, nil
	}
	if os.IsNotExist(err) {
		// 文件或目录不存在
		return false, nil
	}

	return false, err
}
