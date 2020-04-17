package utils

import(
	"os"
)

//是否为目录
func IsDir(path string) bool {
	file, err := os.Open(path)

	if err != nil {
		return false
	}
	defer file.Close()

	stat, err := file.Stat()
	if err != nil {
		return false
	}

	if !stat.IsDir() {
		return false
	}

	return true
}

//获取文件的除去路径名的文件名
func GetFileName(path string) string {
	for i := len(path) - 1; i >= 0; i-- {
		if path[i] == '/' {
			path = path[i+1:]
			break
		}
	}
	return path
}
