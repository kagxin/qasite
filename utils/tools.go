package utils

import (
	"crypto/md5"
	"encoding/hex"
)

// Md5 adsf
func Md5(str string) string {
	h := md5.New()
	h.Write([]byte(str))
	return hex.EncodeToString(h.Sum(nil))
}
