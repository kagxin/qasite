package utils

import (
	"crypto/md5"
	"encoding/hex"
	"strconv"
	"time"
)

// Md5 adsf
func Md5(str string) string {
	h := md5.New()
	h.Write([]byte(str))
	return hex.EncodeToString(h.Sum(nil))
}

// genToken 生成token
func GenToken(userId uint) string {
	h := md5.New()
	h.Write([]byte(strconv.FormatUint(uint64(userId), 10)))
	h.Write([]byte(strconv.FormatUint(uint64(time.Now().UnixNano()), 10)))
	return hex.EncodeToString(h.Sum(nil))
}
