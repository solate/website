package encrypt

import (
	"io"
	"fmt"
	"crypto/md5"
	"crypto/sha1"
)

/**
加密工具方法,进行SHA1加密
 */
func GetSHA1(data string) string {
	t := sha1.New();
	io.WriteString(t,data);
	return fmt.Sprintf("%x",t.Sum(nil));
}



/**
加密工具方法,进行MD5加密
 */
func GetMD5(data string) string {
	t := md5.New();
	io.WriteString(t,data);
	return fmt.Sprintf("%x",t.Sum(nil));

}
