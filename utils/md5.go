package utils

import (
	"crypto/md5"
	"encoding/hex"
)

// MD5V 密码加密
func MD5V(v []byte) string {
	h := md5.New()
	h.Write(v)
	return hex.EncodeToString(h.Sum(v))
}
