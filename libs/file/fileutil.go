package file

import (
	"os"
	"log"
	"strings"
	"path/filepath"
)

//获取当前文件路径
func GetCurrentDirectory() string {
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		log.Fatalln(err)
		os.Exit(1)
	}
	return strings.Replace(dir, "\\", "/", -1)
}