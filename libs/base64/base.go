package base64

import "encoding/base64"

//base64 编码
func Base64Encode(value string) string {
	encode := base64.StdEncoding.EncodeToString([]byte(value))
	return encode
}

//base64解码
func Base64Decde(value string) ([]byte, error) {
	decode, err := base64.StdEncoding.DecodeString(value)
	return decode, err
}

